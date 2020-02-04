@darwin @linux @windows
Scenario: CRC delete
    When executing "crc delete -f" succeeds
    Then stdout should contain "Deleted the OpenShift cluster"
