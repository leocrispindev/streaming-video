package com.br.watcher.model

import org.apache.kafka.common.protocol.types.Field.Str

data class Message(val videoName : String, val videoId : Int, val videoExtension : String, val session : String)
