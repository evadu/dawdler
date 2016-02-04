var tVo;
var Template = {
    name: undefined,
    tVo: undefined,
    status: 1,//1为列表，2为card
    getInstance: function () {
        if (this.tVo == undefined) {
            this.tVo = Object.create(Template);
        }
        return this.tVo;
    },
    initView: function () {//初始化界面
        console.log("去渲染页面，是HTML还是JS了了了了了");
    },
    addEvent: function () {//添加事件
        console.log("给渲染的HTML绑定相应的事件");
        dawdlerEls = DawdlerHTML.getByTag("dawdler");
        for (var i = 0; i < dawdlerEls.length; i++) {
            el = dawdlerEls[i];
            DawdlerEvt.initClick(el);//点击双击事件
            DawdlerEvt.initMouseEvt(el);//鼠标5个事件
        }
    },
    release: function () {//发布出去
        console.log("当发布出去后会生成相应的HTML，模块的分离");
    }
}