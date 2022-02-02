# Notes from the Prometheus up & running book

## Rules

### Recording rules

- They are defined in rules files, using YAML format. You can specify the location of the rules file in your `prometheus.yml` using the `rule_files` top level field.

- Rule files contain zero or more groups of rules. Each group contains 1+ rules.
  
- `expr` is the PromQL expression to be evaluated and output into a metric name, that is specified by `record`.

Example of a rule file (taken from official [Prometheus docs](https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/)):

```yaml
groups:
  - name: example
    rules:
    - record: job:http_inprogress_requests:sum
      expr: sum by (job) (http_inprogress_requests)
```

- Rules in a group will be run sequentially; a first rule whitin a group is evaluated and its output is ingested in a time series database, before the second rule runs. But, groups will run in different times, which helps to spread the load.

#### When to use recording rules

The main use is to make your queries more efficient and reduce cardinality. For example, if for a specific query the number of targets grows exponentially, it would be a good idea to precompute the query value using a rule group. Then you need to fetch the result time series only once - because now you have a recording rule that precomputes the value!

Ideally, you should put all rules for one job in the same group.

You can also aggregate rules, but they should be in a right order when defined in a rule group. For a rule group Y, with 2 rules, A and B, given that rule B uses rule A, rule A must be defined first within Y, since A and B are evaluated sequentially.

Recording rules can be used also for composing range vector functions. Range vector functions by definition cannot be used on the output of functions that produce instant vectors. You cannot use this technique for `rate`,`irate` or `increase`.

## Alerting

Alert flow: defining alerting rules; AlertManager converts alerts into notifications (e.g. emails, pages, chat messages), connecting them to integrations (e.g. PagerDuty).

In Prometheus you define the alert logic; AlertManager picks up the alerts that are firing in Prometheus. AlertManager groups the alerts together and routes them to the correct configured integrations.

The PromQL queries should return results only for the conditions you want to alert on. The alertmanager(s) can be configured in the `prometheus.yml` file. The `rules.yml` file contains the definitions of the alerting rules (and may have also recording rules there).
In `alertmanager.yml` you can have specific alertmanager configuration (e.g. email/notifications).

## Alerting rules

- Can be placed in the same rule group as recording rules.
- Differently from recording rules, filtering is essential to alerting rules.
- Alerts will populate the `METRICS` metric. Besides the labels you configured, an `alertstate` label is added and its value can be `firing` or `pending`.

### for

Until the for condition is met, the alert is considered `pending`. After, the `alertstate` label changes to `firing`.
To check the state of the alerts, go to the `/alerts` endpoint.
