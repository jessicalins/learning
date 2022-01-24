# Notes from the Prometheus up & running book

## Alerting

Alert flow: defining alerting rules; AlertManager converts alerts into notifications (e.g. emails, pages, chat messages), connecting them to integrations (e.g. PagerDuty). In Prometheus you define the alert logic; in AlertManager picks up the alerts that are firing in Prometheus. AlertManager groups the alerts together and routes them to the correct configured integrations.
The PromQL queries should return results only for the conditions you want to alert on. The alertmanager(s) can be configured in the `prometheus.yml` file. The `rules.yml` file contains the definitions of the alerting rules (and may have also recording rules there).
In `alertmanager.yml` you can have specific alertmanager configuration (e.g. email/notifications).
