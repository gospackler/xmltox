import marionette
import base64

client = marionette.Marionette('localhost', port=2828)
client.start_session()
client.navigate("http://direct-stg.hart.com/xml_snap/e1_inpatient.xml")
res = client.screenshot()
print res
##
"""
decRes = base64.b64decode(res)
fh = open("abc.png", "w")
fh.write(decRes)
fh.close()
"""
