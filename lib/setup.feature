@linux
Scenario: CRC setup on Linux
    When executing "crc setup" succeeds
    Then stdout should contain "Caching oc binary"
    And stdout should contain "Checking if CRC bundle is cached in '$HOME/.crc'"
    And stdout should contain "Checking if running as non-root"
    And stdout should contain "Checking if Virtualization is enabled"
    And stdout should contain "Checking if KVM is enabled"
    And stdout should contain "Checking if libvirt is installed"
    And stdout should contain "Checking if user is part of libvirt group"
    And stdout should contain "Checking if libvirt is enabled"
    And stdout should contain "Checking if libvirt daemon is running"
    And stdout should contain "Checking if a supported libvirt version is installed"
    And stdout should contain "Checking for obsolete crc-driver-libvirt"
    And stdout should contain "Checking if libvirt 'crc' network is available"
    And stdout should contain "Checking if libvirt 'crc' network is active"
    And stdout should contain "Checking if NetworkManager is installed"
    And stdout should contain "Checking if NetworkManager service is running"
    And stdout should contain "Checking if /etc/NetworkManager/conf.d/crc-nm-dnsmasq.conf exists"
    And stdout should contain "Writing Network Manager config for crc"
    And stdout should contain "Will use root access: write NetworkManager config in /etc/NetworkManager/conf.d/crc-nm-dnsmasq.conf"
    And stdout should contain "Will use root access: execute systemctl daemon-reload command"
    And stdout should contain "Will use root access: execute systemctl stop/start command"
    And stdout should contain "Checking if /etc/NetworkManager/dnsmasq.d/crc.conf exists"
    And stdout should contain "Writing dnsmasq config for crc"
    And stdout should contain "Will use root access: write dnsmasq configuration in /etc/NetworkManager/dnsmasq.d/crc.conf"
    And stdout should contain "Will use root access: execute systemctl daemon-reload command"
    And stdout should contain "Will use root access: execute systemctl stop/start command"
    And stdout should contain "Setup is complete, you can now run 'crc start -b $bundlename' to start the OpenShift cluster" if bundle is not embedded
    And stdout should contain "Setup is complete, you can now run 'crc start' to start the OpenShift cluster" if bundle is embedded

@darwin
Scenario: CRC setup on Mac
    When executing "crc setup" succeeds
    Then stdout should contain "Caching oc binary"
    And stdout should contain "Checking if running as non-root"
    And stdout should contain "Checking if HyperKit is installed"
    And stdout should contain "Checking if crc-driver-hyperkit is installed"
    And stdout should contain "Installing crc-machine-hyperkit"
    And stdout should contain "Will use root access: change ownership"
    And stdout should contain "Will use root access: set suid"
    And stdout should contain "Checking file permissions"

@windows
Scenario: CRC setup on Windows
    When executing "crc setup" succeeds
    Then stdout should contain "Caching oc binary"
    Then stdout should contain "Unpacking bundle from the CRC binary" if bundle is embedded
    Then stdout should contain "Checking Windows 10 release"
    Then stdout should contain "Checking if Hyper-V is installed"
    Then stdout should contain "Checking if user is a member of the Hyper-V Administrators group"
    Then stdout should contain "Checking if the Hyper-V virtual switch exist"
    
