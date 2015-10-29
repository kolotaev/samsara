(defproject samsara/samsara-core "0.4.0"
  :description "Event stream processing pipeline"

  :url "https://samsara.github.com/"
  :license {:name "Apache License 2.0"
            :url "http://www.apache.org/licenses/LICENSE-2.0"}

  :java-source-paths ["java"]

  :dependencies [[org.clojure/clojure "1.7.0"]
                 [org.slf4j/slf4j-log4j12 "1.7.12"]
                 [org.apache.samza/samza-api        "0.9.1"]
                 [org.apache.samza/samza-kafka_2.10 "0.9.1"]
                 [samsara/moebius "0.3.0"]
                 [samsara/samsara-utils "0.4.0"]
                 [samsara/trackit "0.2.2"]
                 [com.taoensso/timbre "3.4.0"]
                 [org.clojure/tools.cli "0.3.1"]
                 [digest "1.4.4"]
                 [org.clojure/tools.nrepl "0.2.11"]
                 [clj-kafka "0.3.2"]
                 ]

  :main samsara-core.main

  :profiles {:uberjar {:aot :all}
             :dev {:dependencies [[midje "1.7.0"]]
                   :plugins [[lein-midje "3.1.3"]
                             [lein-bin "0.3.5"]]}}

  :jvm-opts ["-server" "-Dfile.encoding=utf-8"]
  :bin {:name "samsara-core" :bootclasspath false})
