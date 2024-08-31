# read input from freq -c, check that the number of c's is approximately 10 times
# the number of d's.
BEGIN {
	nc = 0
	nd = 0
		# number of c's and d's
	epsilon = 0.10
		# acceptable error between nc and 10*nd
}
/^c/ { nc = $2 }
/^d/ { nd = $2 }
END {
	if (nc == 0) {
		print "stats.awk: no c's"
		exit 1
	} else if (nd == 0) {
		print "stats.awk: no d's"
		exit 1
	}
	diff = percdiff(nc, 10*nd)
	if (diff > epsilon) {
		printf "difference between the number of c's (%d) and 10 times the number of d's (10*%d=%d)is too great (%.1f%% > %.1f%%)\n",
			nc, nd, 10*nd, diff*100, epsilon*100
		exit 1
	}
}

# percdiff: percent difference between a and b
function percdiff(a, b) {
	return abs(a - b) / avg(a, b)
}

function abs(x) {
	if (x < 0)
		return -x
	return x
}

function avg(a, b) {
	return (a + b) / 2
}