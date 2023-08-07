package gen

//Keep in mind that some symbols take >1 byte and if used to generate random string, the result may not be human-readable.
//
//åäö - this string, for example, takes 6 bytes, not 3. So if you try to take random bytes from here,
//you will possibly get something like `�å`
//
//Use tools like this one to check https://mothereff.in/byte-counter
const (
	DigitsAndEnglish = `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`
	DigitsAndSwedish = `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖabcdefghijklmnopqrstuvwxyzåäö`
	DigitsAndGerman  = `0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZÄÖÜẞabcdefghijklmnopqrstuvwxyzäöüß`
)
