# Golang Stackdriver format package

## Summary
This package helps generating Google stackdriver conform logging data. 
Currently the scope is limited to generate log messages to print towards STDOUT, so strongly depends on goolge-fluentd agent (pre-installed on GCE or GKE if logging has been activated)

## Installation
    go get -u github.com/yuenfaytsang/stackdriver

## Usage
Import into your Project:


    import (
        ...
        "github.com/yuenfaytsang/golang-stackdriver/log"
        ...
    )
    
Golang Stackdriver Functions:

    # LOG A DEFAULT MESSAGE TO STDOUT (STACKDRIVER CODE 0)
    log.DEFAULT("YOUR DEFAULT MESSAGE HERE").Stdout()
    
    # LOG A DEBUG MESSAGE TO STDOUT (STACKDRIVER CODE 100)
    log.DEBUG("YOUR DEBUG MESSAGE HERE").Stdout()
    
    # LOG A INFO MESSAGE TO STDOUT (STACKDRIVER CODE 200)
    log.INFO("YOUR INFO MESSAGE HERE").Stdout()
        
    # LOG A NOTICE MESSAGE TO STDOUT (STACKDRIVER CODE 300)
    log.NOTICE("YOUR NOTICE MESSAGE HERE").Stdout()
     
    # LOG A WARNING MESSAGE TO STDOUT (STACKDRIVER CODE 400)
    log.WARNING("YOUR WARNING MESSAGE HERE").Stdout()
    
    # LOG A ERROR MESSAGE TO STDOUT (STACKDRIVER CODE 500)
    log.ERROR("YOUR ERROR MESSAGE HERE").Stdout()
        
    # LOG A CRITICAL MESSAGE TO STDOUT (STACKDRIVER CODE 600)
    log.CRITICAL("YOUR CRITICAL MESSAGE HERE").Stdout()
        
    # LOG A ALERT MESSAGE TO STDOUT (STACKDRIVER CODE 700)
    log.ALERT("YOUR ALERT MESSAGE HERE").Stdout()
    
    # LOG A EMERGENCY MESSAGE TO STDOUT (STACKDRIVER CODE 800)
    log.EMERGENCY("YOUR EMERGENCY MESSAGE HERE").Stdout()