import 'package:flutter/material.dart';

import '../../screens/client_page.dart';
import '../../screens/dashboard_page.dart';

class NavItem {
  Icon icon;
  String label;
  Widget page;

  NavItem({this.icon, this.label, this.page});
}

var navItems = [
  NavItem(
    icon: Icon(Icons.home),
    label: 'Home',
    page: DashboardPage(),
  ),
  NavItem(
    icon: Icon(Icons.settings_sharp),
    label: 'Clients',
    page: ClientPage(),
  ),
];
