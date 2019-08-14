function check(document) {
    let name = document.getElementById("name").value;
    let email = document.getElementById("email").value;
    let password = document.getElementById("password").value;
    let group = document.getElementById("group").value;
    if(name == ""){
        document.alert("姓名不能为空");
    }
    if(email == ""){
        document.alert("Email不能为空");
    }
    if(password == ""){
        document.alert("密码不能为空");
    }
    if(group == ""){
        document.alert("你是哪个团队的？");
    }
    if(name != "" && email != "" && password != "" && group != ""){

    }

}