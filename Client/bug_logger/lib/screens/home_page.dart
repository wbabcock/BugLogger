import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../fragments/app_bar/app_bar_gradient.dart';
import '../fragments/bottom_navigation/navigation_items.dart';
import '../fragments/bottom_navigation/navigation_provider.dart';
import '../globals/theme.dart';

class HomePage extends StatefulWidget {
  final String title;

  HomePage({Key key, this.title}) : super(key: key);

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    var provider = Provider.of<BottomNavigationBarProvider>(context);
    return Scaffold(
      appBar: AppBar(
        title: Text(this.widget.title),
        flexibleSpace: AppBarGradient(),
      ),
      body: Container(
        decoration: BoxDecoration(
          gradient: RadialGradient(
            center: Alignment.topRight,
            radius: 0.75,
            colors: [
              bgPurple,
              bgBlack,
            ],
          ),
        ),
        height: MediaQuery.of(context).size.height,
        padding: EdgeInsets.symmetric(horizontal: 12.0, vertical: 15.0),
        child: navItems[provider.currentIndex].page,
      ),
      bottomNavigationBar: BottomNavigationBar(
        currentIndex: provider.currentIndex,
        backgroundColor: bgBlack,
        selectedItemColor: accentDark,
        onTap: (index) {
          provider.currentIndex = index;
        },
        items: navItems
            .map(
              (item) => BottomNavigationBarItem(
                icon: item.icon,
                label: item.label,
                backgroundColor: bgBlack,
              ),
            )
            .toList(),
      ),
    );
  }
}
