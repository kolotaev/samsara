(ns ingestion-api.system
  (:require [com.stuartsierra.component :as component]
            [reloaded.repl :refer [system init start stop go reset]]
            [ingestion-api.components.backend :as backend]
            [ingestion-api.components.http-server :as http]
            [ingestion-api.components.mqtt-server :as mqtt]
            [ingestion-api.components.admin-server :as admin]))


(defn ingestion-api-system
  [config]
  (component/system-map
   :backend     (backend/new-backend config)
   :admin-server (admin/new-admin-server config)
   :http-server (component/using
                 (http/new-http-server config) {:backend :backend})
   :mqtt-server (component/using
                 (mqtt/new-mqtt-server config) {:backend :backend})))
