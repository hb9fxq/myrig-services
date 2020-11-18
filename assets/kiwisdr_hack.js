<script async src="https://www.googletagmanager.com/gtag/js?id=UA-10245752-8"></script>



<!-- Global site tag (gtag.js) - Google Analytics -->

<script>
    window.dataLayer = window.dataLayer || [];
    function gtag(){dataLayer.push(arguments);}
    gtag('js', new Date());

    gtag('config', 'UA-10245752-8');

    function updateAntennaPos(){

    fetch("https://schwarzenburg.myrig.ch/myrig-services/rotor")
        // bald, ja: fetch("https://schwarzenburg.myrig.ch/myrig-services/kiwi")
        .then(function (response) {
            return response.json();
        })
        .then(function (myJson) {
            document.getElementById('deg').textContent = 'Connected Antenna: Optibeam@' + myJson.Optibeam.Deg + '°'
            // bald, ja: document.getElementById('deg').textContent = 'Connected Antenna: ' + myJson.HeaderText;
            console.log(myJson.Deg);
        })
        .catch(function (error) {
            console.log("Error: " + error);
        });

}

</script>
<script>

    function myFunction(){
    console.log('update rotor Called')
    updateAntennaPos()
}

    myFunction();

    setInterval(function(){
    myFunction()
}, 15000)


    document.addEventListener("DOMContentLoaded", function(event) {

});

</script>


