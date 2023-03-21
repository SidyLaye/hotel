import 'package:flutter/material.dart';

import 'components/body.dart';

class HotelFormScreen extends StatelessWidget {
  static String routeName = "/hotel_form";
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Information sur l'hotel"),
      ),
      body: Body(),
    );
  }
}
