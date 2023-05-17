package service.database

import java.sql.Connection
import java.sql.DriverManager

object Database {
    private lateinit var conn : Connection
    init {
      try {
          conn = DriverManager.getConnection("")
          println("Connection ok")

      }catch (e : Exception) {
        throw e;
      }
    }

    public fun getConn(): Connection {
        return conn;
    }
}