@darwin @linux @windows
Scenario: CRC forcible stop
    When executing "crc stop -f"
    Then stdout should match "(.*)[Ss]topped the OpenShift cluster"
