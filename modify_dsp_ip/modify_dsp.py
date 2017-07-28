#!/usr/bin/python
import redis
import json
dsp_list = [12, 33, 48, 5, 63, 60, 17, 12, 37]
redis_host = "127.0.0.1"
redis_port = 1025

r = redis.Redis(host=redis_host, port=redis_port,db=0)
def modify_dsp(dsp, pOld):
  dsp = str(dsp)
  dsp_txt = r.get("dsp."+dsp)
  old_dsp_txt = dsp_txt
  dsp_json = json.loads(dsp_txt)
  bid_url =  dsp_json["bid_url"]
  bid_url_list = bid_url.split("/")

  if len(dsp) < 2:
    bid_url_list[2] = "127.0.0.1:"+"3800"+dsp
  else:
    bid_url_list[2] = "127.0.0.1:"+"380"+dsp
  bid_url = "/".join(bid_url_list)
  dsp_json["bid_url"] = bid_url

  dsp_txt = json.dumps(dsp_json)

  if pOld:
    print "#redis-cli -h {HOST} -p {PORT} set dsp.{DSPID} '{DSP_TXT}'".format (
      HOST=redis_host,
      PORT=redis_port,
      DSPID=dsp,
      DSP_TXT=old_dsp_txt)
  else:
    print "redis-cli -h {HOST} -p {PORT} set dsp.{DSPID} '{DSP_TXT}'".format (
      HOST=redis_host,
      PORT=redis_port,
      DSPID=dsp,
      DSP_TXT=dsp_txt)
  
if __name__=='__main__':
  for dsp in dsp_list:
    modify_dsp(dsp,True)
  for dsp in dsp_list:
    modify_dsp(dsp,False)
  print "redis-cli -h {HOST} -p {PORT} save".format (
      HOST=redis_host,
      PORT=redis_port)
  



