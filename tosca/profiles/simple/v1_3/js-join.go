// This file was auto-generated from a YAML file

package v1_3

func init() {
	Profile["/tosca/simple/1.3/js/join.js"] = `

// [TOSCA-Simple-Profile-YAML-v1.3] @ 4.3.2

function evaluate() {
	var length = arguments.length;
	if ((length < 1) || (length > 2))
		throw 'must have 1 or 2 arguments';
	var delimiter = (length == 2) ? arguments[1] : '';
	var args = arguments[0];
	length = args.length;
	var a = [];
	for (var i = 0; i < length; i++) {
		var argument = args[i];
		if (argument.$string !== undefined)
			argument = argument.$string;
		a.push(argument);
	}
	return a.join(delimiter);
}
`
}
