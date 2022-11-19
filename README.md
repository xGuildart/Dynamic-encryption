<h1>Dynamic encryption</h1>
<p>The main purpose for this library is to encrypt strings with RSA encryption and make it not dependant to public/private key only, but the application that use the RSA algorithm too<p>
<p>Which mean even if you have the key, you can't encrypt/decrypt the string using RSA (it's more useful when encrypting files)</p>

RSA is simply the algorithm of Rijndael, by using the polynom 3x^3+x^2+x+2 and it's inverse 11x^3+13x^2+9x+14 to encrypt/decrypt respectively, and making changes on both (key, string) using substraction bytes, shifting rows and mixing columns...

In other hand; he uses one of the lowest polynom power that it's Matrix verify M * M^-1 = I , in vectorial space Z(16)4: (16: Hexadecimal, 4 for matrix range = polynom degree (power)), 
[why he chooses that, you can read his document on the web or just wikipedia; it's related to length of public key, universal length of trams ...]
Which give:

3 1 1 2     *      11 13  9 14    =   1 0 0 0 </br>
2 3 1 1     *      14 11 13  9    =   0 1 0 0 </br>
1 2 3 1     *       9 14 11 13    =   0 0 1 0 </br>
1 1 2 3     *      13  9 14 11    =   0 0 0 1 </br>


So to quit the habit of RSA; I simply, decide to change the couple (M,M^-1) and I can get another publicKey/privateKey that solve the problem with our form of RSA

This implementation contain 3 mode of using RSA 

<h4>Default</h4>
Is The same as RSA-ISO
<h4>Special</h4>
RSA using our special polynom: x^3+x^2+4x+5
<h4>Dynamic</h4>
RSA using a generated Matrix based on the given publicKey
It adds another tier of complexity to Rijndael Algorithm, and make the solving rely on this package....

You can add/change/mix some spices and get different result, if you intend to integrate it on your work... 
