CREATE TABLE `metadata` (
  `id` int(11) NOT NULL,
  `metric_id` int(11) NOT NULL,
  `key_for_value` varchar(255) NOT NULL,
  `value` varchar(4000) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `metrics` (
  `id` int(11) NOT NULL,
  `key_for_value` varchar(255) NOT NULL,
  `value` varchar(4000) NOT NULL,
  `timestamp` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


ALTER TABLE `metadata`
  ADD PRIMARY KEY (`id`),
  ADD KEY `metric_id` (`metric_id`);

ALTER TABLE `metrics`
  ADD PRIMARY KEY (`id`);


ALTER TABLE `metadata`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
ALTER TABLE `metrics`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

ALTER TABLE `metadata`
  ADD CONSTRAINT `metadata_metrics_FK` FOREIGN KEY (`metric_id`) REFERENCES `metrics` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
