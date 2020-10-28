import 'package:flutter/material.dart';

import '../../globals/theme.dart';

class AppBarGradient extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        gradient: LinearGradient(
          colors: [
            bgBlack,
            bgPurple,
          ],
          stops: [
            0.0,
            0.9,
          ],
        ),
      ),
    );
  }
}
