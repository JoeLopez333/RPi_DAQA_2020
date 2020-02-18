import json 
import os

dirList = os.listdir('./')

timelist = []

for i in dirList:
	basename, extension = os.path.splitext(i)
	#print(extension)
	if extension == '.json':
		with open(str(i), 'r') as f:
			#print("opened + " + str(i))
			line = f.readline()
			cnt = 0
			while line:
				strlin =str(line)[:-1]
				#print(strlin)	
				try: jsonobj = json.loads(strlin)
				except:
					print(cnt)
					break
				
				if jsonobj['msg'][4:6] == "d1":
					timelist.append(jsonobj['time'])		
						
				try: line = f.readline()
				except: 
					print("End of file")
					line = NULL
				cnt+=1
			print(cnt)
#print(timelist[1])
#time = int.from_bytes(timelist[0],byteorder = 'big')
#time = timelist[0].decode('ascii', 'ignore') 
#print(type(timelist[0]))
#
