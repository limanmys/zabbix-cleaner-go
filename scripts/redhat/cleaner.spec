Name: zabbix-cleaner
Version: %VERSION%
Release: 0
License: MIT
Prefix: /opt
Summary: Zabbix Database Cleaner
Group: Applications/System
BuildArch: x86_64

%description
Zabbix Database Cleaner

%pre

%prep

%build

%install
cp -rfa %{_app_dir} %{buildroot}

%post -p /bin/bash
chown -R root:root /opt/zabbix-cleaner
chmod -R 770 /opt/zabbix-cleaner
if [ -f "/etc/systemd/system/zabbix-cleaner.service" ]; then
    rm -rf /etc/systemd/system/zabbix-cleaner.service
    systemctl disable zabbix-cleaner.service
    systemctl stop zabbix-cleaner.service
    systemctl daemon-reload
fi
echo """
[Unit]
Description=Zabbix Database Cleaner

[Service]
WorkingDirectory=/opt/zabbix-cleaner
Type=simple
Restart=always
RestartSec=5s
ExecStart=/opt/zabbix-cleaner/main

[Install]
WantedBy=multi-user.target
""" > /usr/lib/systemd/system/zabbix-cleaner.service

systemctl daemon-reload
systemctl enable zabbix-cleaner.service
systemctl restart zabbix-cleaner.service

%clean

%files
%defattr(0770, root, root)
/opt/zabbix-cleaner/*
/opt/zabbix-cleaner/.env.example

%define _unpackaged_files_terminate_build 0