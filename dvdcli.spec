%define        _topdir  %{rpmbuild}
%define        _tmppath %{_topdir}/tmp

Summary: Tool for managing remote & local storage.
Name: dvdcli
Version: %{v_semver}
Release: 1
License: Apache License
Group: Applications/Storage
#Source: https://github.com/emccode/dvdcli/archive/master.zip
URL: https://github.com/emccode/dvdcli
Vendor: EMC{code}
Packager: Andrew Kutz <sakutz@gmail.com>
BuildArch: %{v_arch}
BuildRoot: %{_tmppath}/%{name}-%{version}-%{release}

%description
A guest based storage introspection tool that
allows local visibility and management from cloud
and storage platforms.

%prep

%build

%install
install -D %{dvdcli} $RPM_BUILD_ROOT/usr/bin/dvdcli

%clean
#rm -rf "$RPM_BUILD_ROOT"

%files
%attr(0755, root, root) /usr/bin/dvdcli
