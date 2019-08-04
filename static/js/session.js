function session() {
    if ((Data["Username"])==null) {
        $("#a").style.display = 'block';
        $("#b").style.display = 'none';
    } else {
        $("#a").style.display = 'none';
        $("#b").style.display = 'block';
    }



}