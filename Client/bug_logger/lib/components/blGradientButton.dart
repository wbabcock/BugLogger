import 'package:bug_logger/globals/theme.dart';
import 'package:flutter/material.dart';

class BLGradientButton extends StatelessWidget {
  final Function onPressed;
  final Widget child;

  BLGradientButton({Key key, this.onPressed, this.child}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: this.onPressed,
      child: Container(
        margin: EdgeInsets.only(
          top: 15,
          bottom: 15,
        ),
        padding: EdgeInsets.symmetric(
          vertical: 10,
          horizontal: 20,
        ),
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(30),
          gradient: LinearGradient(
            colors: [
              accentBright,
              accentDark,
            ],
            begin: Alignment.topRight,
            end: Alignment.bottomLeft,
          ),
          boxShadow: [
            BoxShadow(
              color: accentBright,
              offset: Offset(0, 5),
              blurRadius: 20,
              spreadRadius: 0,
            ),
          ],
        ),
        child: Center(child: this.child),
      ),
    );
  }
}
