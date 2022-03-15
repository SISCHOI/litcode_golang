func replaceSpace(s string) string {
	var snew string
	for _, i := range s {
		switch i { //注意此处i取出来的应该都是char
		case ' ':
			{
				snew += "%20"
			}

		default:
			snew += string(i)
		}

	}
	return snew

}