package com.br.watcher.broker

import com.typesafe.config.ConfigException.Null
import org.apache.kafka.clients.admin.KafkaAdminClient
import org.apache.kafka.common.KafkaException

object BrokerAdmin {

    // TODO define admin properties
    init {

    }

    // TODO create topic if not exists
    fun checkTopic() {
        try {



        }catch (e : KafkaException) {
            throw e
        }
    }

}