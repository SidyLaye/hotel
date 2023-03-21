import 'package:flutter/material.dart';
import 'package:shop_app/components/default_button.dart';

class Link {
  String title;
  MaterialPageRoute? direction;
  Link({required this.title, this.direction});
}

class RenderButton extends StatefulWidget {
  const RenderButton({key, required this.link, required this.title});

  final List<Link> link;
  final String title;

  @override
  State<RenderButton> createState() => _RenderButton();
}

class _RenderButton extends State<RenderButton> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Center(
        child: Expanded(
            child: ListView.builder(
          itemCount: widget.link.length,
          itemBuilder: (BuildContext context, int index) {
            return Container(
                margin: const EdgeInsets.only(top: 20.0),
                child: DefaultButton(
                    press: () {
                      if (widget.link[index].direction != null)
                        Navigator.push(context,
                            widget.link[index].direction as Route<Object?>);
                    }, // This child can be everything. I want to choose a beautiful Text Widget
                    text: widget.link[index].title));
          },
        )),
      ),
    );
  }
}
