$(function (){
    app.init();
})

var app={
    init:function(){
        this.getCaptcha()
        this.captchaImgChange()
    },
    getCaptcha:function(){
        $.get("/admin/captcha?t="+Math.random(),function(response){
            console.log(response)
            $("#captchaId").val(response.captchaId)
            $("#captchaImg").attr("src",response.captchaImage)
        })
    },
    captchaImgChange:function(){
        const that = this;
        $("#captchaImg").click(function(){
            that.getCaptcha()
        })
    }
}