import 'package:flutter/material.dart';

class ClientPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.stretch,
      children: [
        RaisedButton(
          onPressed: () => {
            showModalBottomSheet(
                context: context,
                builder: (BuildContext bc) {
                  return Container();
                }),
          },
          child: Text(
            "Add Client",
          ),
        ),
      ],
    );
  }
}
