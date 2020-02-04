@darwin @linux @windows
Scenario: Start CRC
    When starting CRC with default bundle and default hypervisor succeeds
    Then stdout should contain "Started the OpenShift cluster"
