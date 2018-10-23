#########################################################################
# File Name: run.sh
# Author: shanhai2015
#  
# Created Time: 2018年10月22日 星期一 11时24分05秒
#########################################################################
#!/bin/bash
nohup ./smsEmailServices >> /work/logs/smsEmailServices/smsEmailServices.log 2>&1 &
