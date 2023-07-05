package com.br.watcher.model

import org.apache.kafka.common.protocol.types.Field.Str

data class Message(var videoName : String, var videoId : Int, var videoExtension : String)
