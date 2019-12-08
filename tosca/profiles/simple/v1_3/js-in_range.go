// This file was auto-generated from a YAML file

package v1_3

func init() {
	Profile["/tosca/simple/1.3/js/in_range.js"] = `

// [TOSCA-Simple-Profile-YAML-v1.3] @ 3.6.3
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.6.3
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.5.2

clout.exec('tosca.helpers');

function validate(v, lower, upper) {
	if (arguments.length !== 3)
		throw 'must have 2 arguments';
	v = tosca.getComparable(v);
	return (v >= tosca.getComparable(lower)) && (v <= tosca.getComparable(upper));
}
`
}
