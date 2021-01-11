

function getValueDel(obj) {
    var command = $(obj).attr("id");
    // console.log(command);
    $('#inp1').val(command);
    jQuery('#inp2').removeAttr('disabled').attr('value',"请添加参数，如遇多参，请用空格分开。");
}

function getValueAdd(obj) {
    var command = $(obj).attr("id");
    // console.log(command);
    $('#inp1').val(command);
    jQuery('#inp2').attr('disabled', 'disabled').removeAttr('value');
    // jQuery('#inp2').removeAttr('value')

}

function callMe() {
    alert("电话：17692111675，QQ：1420014281")
}

function getCommand() {
    var command = $(obj).attr("command");
    console.log(command);
    $('#inp1').val(this.name)
}