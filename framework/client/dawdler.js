//HTML的总标记
var Dawdler = {
    getInstance: function () {
        dawdler = DawdlerHTML.getByID("dawdler_global_main");
        if (dawdler == undefined) {
            document.write('<dawdler id="dawdler_global_main" style="display: none"></dawdler>')
        }
        return DawdlerHTML.getByID("dawdler_global_main");
    }
}
//Dawdler 常见获取HTML的属性参数=.
var DawdlerHTML = {
    getAttr: function (el, name) {
        att = el.attributes[name]
        return att == undefined ? undefined : att.value;
    },
    setAtt: function (el, name, val) {

    },
    hasAtt: function (el, name) {
        att = el.attributes[name]
        return att != undefined;
    },
    getByID: function (name) {
        return document.getElementById(name);
    },
    getByTag: function (name) {
        return document.getElementsByTagName(name);
    }
}
//Dawdler常见事件枚举
var COMM_EVT = {
    CLICK: "dawdler_Click",
    DOUBLE_CLICK: "dawdler_dbClick",
    MOUSE_OVER: "dawdler_mouseOver",
    MOUSE_OUT: "dawdler_mouseOut",
    MOUSE_UP: "dawdler_mouseUp",
    MOUSE_DOWN: "dawdler_mouseDown",
    MOUSE_MOVE: "dawdler_mouseMove",
    getBind: function (evtType) {
        switch (evtType) {
            case COMM_EVT.CLICK:
                return "click";
            case COMM_EVT.MOUSE_DOWN:
                return "mousedown";
            case COMM_EVT.MOUSE_UP:
                return "mouseup";
            case COMM_EVT.MOUSE_OUT:
                return "mouseout";
            case COMM_EVT.MOUSE_OVER:
                return "mouseover";
            case COMM_EVT.MOUSE_MOVE:
                return "mousemove";
        }
    }
}

//Dawdler添加鼠标的一堆事件
var DawdlerEvt = {
    clickIntervalId: 0,
    clickIntervalTime: 200,
    currentEvt: undefined,
    initClick: function (el) {//是否初始化点击事件
        clickFun = DawdlerHTML.getAttr(el, COMM_EVT.CLICK);
        dbClickFun = DawdlerHTML.getAttr(el, COMM_EVT.DOUBLE_CLICK);
        if (clickFun != undefined || dbClickFun != undefined) {//有此属性必加事件
            if (clickFun != undefined && dbClickFun != undefined) {//单双击同时存在
                el.addEventListener(COMM_EVT.getBind(COMM_EVT.CLICK), function (evt) {
                    DawdlerEvt.currentEvt = evt;
                    if (DawdlerEvt.clickIntervalId != 0) {//双击事件
                        DawdlerEvt.doClick(true);
                    } else {//单击事件
                        DawdlerEvt.clickIntervalId = setTimeout("DawdlerEvt.doClick()", DawdlerEvt.clickIntervalTime);
                    }
                })
            } else if (clickFun != undefined) {//只有单击
                el.addEventListener(COMM_EVT.getBind(COMM_EVT.CLICK), function (evt) {
                    DawdlerEvt.currentEvt = evt;
                    DawdlerEvt.doClick();
                });
            } else {//只有双击
                el.addEventListener(COMM_EVT.getBind(COMM_EVT.CLICK), function (evt) {
                    DawdlerEvt.currentEvt = evt;
                    if (DawdlerEvt.clickIntervalId != 0) {//双击事件
                        DawdlerEvt.doClick(true);
                    } else {//单击事件啥也不做
                        DawdlerEvt.clickIntervalId = setTimeout("console.log('只有双击，点一下无效啦!');DawdlerEvt.clickIntervalId=0;", DawdlerEvt.clickIntervalTime)
                    }
                })
            }
        }
    },
    doClick: function (dbClick) {//点击事件处事
        evtType = dbClick ? COMM_EVT.DOUBLE_CLICK : COMM_EVT.CLICK;
        if (dbClick == true) {
            clearTimeout(DawdlerEvt.clickIntervalId);
        }
        el = DawdlerEvt.currentEvt.target
        clickFun = DawdlerHTML.getAttr(el, evtType);
        if (clickFun != "") {//执行默认事件
            eval(clickFun);
        } else {//加载默认NOVA_CLICK事件，可透传出去
            dawdlerEvt = new Event(DawdlerEvt.currentEvt.name);
            dawdlerEvt.initEvent(evtType);
            dawdler = Dawdler.getInstance();
            dawdler.innerHTML = el.outerHTML;
            dawdler.dispatchEvent(dawdlerEvt);
        }
        DawdlerEvt.clickIntervalId = 0;
    },
    initMouseEvt: function (el) {//是否初始化 鼠标的所有事件
        evtList = [COMM_EVT.MOUSE_OVER, COMM_EVT.MOUSE_OUT, COMM_EVT.MOUSE_UP, COMM_EVT.MOUSE_DOWN, COMM_EVT.MOUSE_MOVE]
        for (var i = 0; i < evtList.length; i++) {
            evtType = evtList[i];
            evtFun = DawdlerHTML.getAttr(el, evtType);
            if (evtFun == undefined)continue;
            console.log(el);
            this.loadDawdlerMouseEvt(el,evtType);
        }
    },
    loadDawdlerMouseEvt:function (el,evtType) {//初始化所有的鼠标事件
        switch (evtType) {
            case COMM_EVT.MOUSE_OVER:
                el.addEventListener(COMM_EVT.getBind(evtType), function (evt) {
                    DawdlerEvt.currentEvt = evt;
                    el = DawdlerEvt.currentEvt.target
                    evtFun = DawdlerHTML.getAttr(el, COMM_EVT.MOUSE_OVER);
                    if (evtFun != "") {//执行默认事件
                        eval(evtFun);
                    } else {//加载默认NOVA_CLICK事件，可透传出去
                        DawdlerEvt.doMouseEvt(COMM_EVT.MOUSE_OVER);
                    }
                });
                break;
            case COMM_EVT.MOUSE_OUT:
                el.addEventListener(COMM_EVT.getBind(evtType), function (evt) {
                    DawdlerEvt.currentEvt = evt;
                    el = DawdlerEvt.currentEvt.target
                    evtFun = DawdlerHTML.getAttr(el, COMM_EVT.MOUSE_OUT);
                    if (evtFun != "") {//执行默认事件
                        eval(evtFun);
                    } else {//加载默认NOVA_CLICK事件，可透传出去
                        DawdlerEvt.doMouseEvt(COMM_EVT.MOUSE_OUT);
                    }
                });
                break;
            case COMM_EVT.MOUSE_UP:
                el.addEventListener(COMM_EVT.getBind(evtType), function (evt) {
                    DawdlerEvt.currentEvt = evt;
                    el = DawdlerEvt.currentEvt.target
                    evtFun = DawdlerHTML.getAttr(el, COMM_EVT.MOUSE_UP);
                    if (evtFun != "") {//执行默认事件
                        eval(evtFun);
                    } else {//加载默认NOVA_CLICK事件，可透传出去
                        DawdlerEvt.doMouseEvt(COMM_EVT.MOUSE_UP);
                    }
                });
                break;
            case COMM_EVT.MOUSE_DOWN:
                el.addEventListener(COMM_EVT.getBind(evtType), function (evt) {
                    DawdlerEvt.currentEvt = evt;
                    el = DawdlerEvt.currentEvt.target
                    evtFun = DawdlerHTML.getAttr(el, COMM_EVT.MOUSE_DOWN);
                    if (evtFun != "") {//执行默认事件
                        eval(evtFun);
                    } else {//加载默认NOVA_CLICK事件，可透传出去
                        DawdlerEvt.doMouseEvt(COMM_EVT.MOUSE_DOWN);
                    }
                });
                break;
            case COMM_EVT.MOUSE_MOVE:
                el.addEventListener(COMM_EVT.getBind(evtType), function (evt) {
                    DawdlerEvt.currentEvt = evt;
                    el = DawdlerEvt.currentEvt.target
                    evtFun = DawdlerHTML.getAttr(el, COMM_EVT.MOUSE_MOVE);
                    if (evtFun != "") {
                        eval(evtFun);
                    } else {
                        DawdlerEvt.doMouseEvt(COMM_EVT.MOUSE_MOVE);
                    }
                });
                break;
        }
    },
    doMouseEvt: function (type) {//鼠标事没有时派发处理
        switch (type) {
            case COMM_EVT.MOUSE_OVER:
                break;
            case COMM_EVT.MOUSE_OUT:
                break;
            case COMM_EVT.MOUSE_UP:
                break;
            case COMM_EVT.MOUSE_DOWN:
                break;
            case COMM_EVT.MOUSE_MOVE:
                break;
        }
        dawdlerEvt = new Event(DawdlerEvt.currentEvt.name);
        dawdlerEvt.initEvent(evtType);
        dawdler = Dawdler.getInstance();
        dawdler.innerHTML = el.outerHTML;
        dawdler.dispatchEvent(dawdlerEvt);
    }
}