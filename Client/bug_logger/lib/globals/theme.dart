import 'package:flutter/material.dart';

var bgPurple = Color(0xff231b3a);
var bgBlack = Color(0xff1f2128);
var bgLight1 = Color(0xff292e3c);
var bgLight2 = Color(0xff1d202a);
var bgDark = Color(0xff1f2128);

var accentBright = Color(0xffbe3eff);
var accentDark = Color(0xff7839ff);

/// Main Dark Theme
var darkTheme = ThemeData(
  // default brightness and colors
  brightness: Brightness.dark,
  primarySwatch: Colors.blueGrey,
  accentColor: accentBright,
  visualDensity: VisualDensity.adaptivePlatformDensity,

  // appbar theme
  appBarTheme: AppBarTheme(
    color: bgBlack,
    elevation: 0.0,
  ),

  buttonTheme: ButtonThemeData(
    buttonColor: accentDark,
  ),

  // default text theme
  textTheme: TextTheme(
    headline1: TextStyle(
      fontSize: 32.0,
      fontWeight: FontWeight.bold,
      color: Colors.white54,
    ),
    button: TextStyle(
      fontSize: 12.0,
      color: Colors.white,
    ),
  ),
);
