#!/bin/bash

/opt/splunk/bin/splunk start &&
tail -f /opt/splunk/var/log/splunk/splunkd.log
