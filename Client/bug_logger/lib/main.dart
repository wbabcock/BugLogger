import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'fragments/bottom_navigation/navigation_provider.dart';
import 'globals/theme.dart';
import 'screens/home_page.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Bug Logger',
      theme: darkTheme,
      home: ChangeNotifierProvider<BottomNavigationBarProvider>(
        create: (BuildContext context) => BottomNavigationBarProvider(),
        child: HomePage(title: 'Bug Logger'),
      ),
    );
  }
}
