import syslog

syslog.syslog(syslog.LOG_ERR, 'Sending a log message [LOG_ERR] through syslog_module!')
syslog.syslog(syslog.LOG_EMERG, 'Sending a log message [LOG_EMERG] through syslog_module!')
syslog.syslog(syslog.LOG_CRIT, 'Sending a log message [LOG_CRIT] through syslog_module!')

#LOG_EMERG, LOG_ALERT, LOG_CRIT, LOG_ERR, LOG_WARNING, LOG_NOTICE, LOG_INFO, LOG_DEBUG
#Emergency System is Unusable
#Alert: Action must be taken immediately
#Critical: Critical Conditions
#Error: Error Conditions
#Warning: Warning Conditions
#Notice: Normal but Significant Condition
#Informational: Informational messages
#Debug: Debug-Level messages