var f=(M,p,i)=>new Promise((E,m)=>{var C=r=>{try{b(i.next(r))}catch(D){m(D)}},h=r=>{try{b(i.throw(r))}catch(D){m(D)}},b=r=>r.done?E(r.value):Promise.resolve(r.value).then(C,h);b((i=i.apply(M,p)).next())});import{A as Z,P as ee,j as R,r as k,al as ae,B as g,D as v,w as u,a5 as s,u as a,bo as P,bp as T,bv as q,aa as G,ao as O,a1 as U,ad as c,J as F,bw as V,aS as S,a4 as te,aY as se,H as x,M as ue,br as ne,K as _,bq as w,bs as le}from"./vendor.35b582d5.js";/* empty css               *//* empty css               *//* empty css              *//* empty css              *//* empty css               *//* empty css               */import{a as oe,f as ie,an as re,ao as A,ad as de,ae as pe,ap as ce,aq as ge}from"./index.044d747f.js";const _e=c("\u641C\u7D22"),me=["onClick"],fe={key:0},ve=c("\u542F\u7528"),Ce={key:1},ye=c("\u7981\u7528"),ke={key:0},Se=c("\u54CD\u5E94"),he={key:1},be=c("\u6C89\u9ED8"),De=c("\u542F\u7528"),Ie=c("\u7981\u7528"),Fe=c("\u8FD8\u539F"),xe=c("\u54CD\u5E94"),Le=c("\u6C89\u9ED8"),Re=Z({setup(M){const{prefixCls:p}=oe("plugin");console.log("prefixCls",p);const i=ee({pluginNames:[],groupIdList:[0],friendIdList:[]}),E=R(()=>({labelCol:{span:4},wrapperCol:{span:14}})),m=k([]),C=k([]),h=k([]),b=ie(),r=k([]),D=[{title:"\u63D2\u4EF6\u540D",dataIndex:"name",key:"name",slots:{customRender:"name"}},{title:"\u7FA4\u540D",dataIndex:"group_name",key:"group_name"},{title:"\u597D\u53CB\u540D",dataIndex:"friend_name",key:"friend_name"},{title:"\u63D2\u4EF6\u72B6\u6001",dataIndex:"pluginStatus",key:"pluginStatus",slots:{customRender:"pluginStatus"}},{title:"\u54CD\u5E94\u72B6\u6001",dataIndex:"responseStatus",key:"responseStatus",slots:{customRender:"responseStatus"}},{title:"\u64CD\u4F5C",dataIndex:"action",key:"action",slots:{customRender:"action"}}],L=k(!1),I=k(),z=(n,o)=>f(this,null,function*(){I.value=yield re({groupId:n,name:o}),L.value=!0}),Q=n=>{console.log(n),L.value=!1},H=n=>f(this,null,function*(){m.value=yield A({groupId:n})}),j=R(()=>b.getQQ),J=()=>f(this,null,function*(){C.value=yield de({selfId:j.value}),C.value.unshift({group_id:0,group_name:"\u5168\u90E8\u7FA4\u804A",group_create_time:0,group_level:0,max_member_count:0,member_count:0})}),K=()=>f(this,null,function*(){h.value=yield pe({selfId:j.value})}),B=(n,o,e)=>f(this,null,function*(){yield ce({groupId:n,name:o,status:e}),r.value.filter(l=>l.gid===n&&l.name===o).forEach(l=>{l.pluginStatus=e})}),N=(n,o)=>f(this,null,function*(){yield ge({groupId:n,status:o}),r.value.filter(e=>e.gid===n).forEach(e=>{e.responseStatus=o})}),Y=n=>{console.log(`selected ${n}`)},W=n=>{console.log(`selected ${n}`)},X=()=>f(this,null,function*(){r.value=[];const n=i.groupIdList.filter(t=>t>=0),o=i.groupIdList.filter(t=>t<0).map(t=>-t);let e=C.value.filter(t=>n.includes(t.group_id));for(let t of e){let y=(yield A({groupId:t.group_id})).filter(d=>i.pluginNames.includes(d.name));for(let d of y)r.value.push({key:t.group_name+d.name,name:d.name,friend_name:"",group_name:t.group_name,gid:t.group_id,pluginStatus:d.pluginStatus,responseStatus:d.responseStatus})}let l=h.value.filter(t=>o.includes(t.user_id));for(let t of l){let y=(yield A({groupId:-t.user_id})).filter(d=>i.pluginNames.includes(d.name));for(let d of y)r.value.push({key:t.nickname+d.name,name:d.name,friend_name:t.nickname,group_name:"",gid:-t.user_id,pluginStatus:d.pluginStatus,responseStatus:d.responseStatus})}});return H(0),ae(()=>{J(),K()}),(n,o)=>(g(),v("div",null,[u(a(se),te({layout:"horizontal"},a(E),{model:a(i)}),{default:s(()=>[u(a(P),{label:"\u63D2\u4EF6\u540D"},{default:s(()=>[u(a(T),{value:a(i).pluginNames,"onUpdate:value":o[0]||(o[0]=e=>a(i).pluginNames=e),mode:"multiple",style:{width:"40%"},placeholder:"\u8BF7\u9009\u62E9\u63D2\u4EF6\u540D",options:m.value.map((e,l)=>({value:m.value[l].name,label:m.value[l].name+" ( "+m.value[l].brief+" )"})),onChange:Y},null,8,["value","options"])]),_:1}),u(a(P),{label:"\u7FA4\u804A & \u597D\u53CB"},{default:s(()=>[u(a(T),{value:a(i).groupIdList,"onUpdate:value":o[1]||(o[1]=e=>a(i).groupIdList=e),mode:"multiple",style:{width:"40%"},placeholder:"\u8BF7\u9009\u62E9\u7FA4\u804A & \u597D\u53CB",onChange:W},{default:s(()=>[u(a(q),{label:"\u7FA4\u804A"},{default:s(()=>[(g(!0),v(G,null,O(C.value,e=>(g(),U(a(V),{value:e.group_id,key:e.group_id},{default:s(()=>[c(F(e.group_name+" ("+e.group_id+")"),1)]),_:2},1032,["value"]))),128))]),_:1}),u(a(q),{label:"\u597D\u53CB"},{default:s(()=>[(g(!0),v(G,null,O(h.value,e=>(g(),U(a(V),{value:-e.user_id,key:e.user_id},{default:s(()=>[c(F(e.nickname+" ("+e.user_id+")"),1)]),_:2},1032,["value"]))),128))]),_:1})]),_:1},8,["value"])]),_:1}),u(a(P),{"wrapper-col":{span:14,offset:4}},{default:s(()=>[u(a(S),{type:"primary",onClick:X},{default:s(()=>[_e]),_:1})]),_:1})]),_:1},16,["model"]),u(a(le),{dataSource:r.value,columns:D},{name:s(({record:e})=>{var l;return[x("a",{onClick:t=>z(e.gid,e.name)},F(e.name),9,me),u(a(ue),{visible:L.value,"onUpdate:visible":o[2]||(o[2]=t=>L.value=t),title:(l=I.value)==null?void 0:l.name,onOk:Q},{default:s(()=>{var t,$,y;return[x("span",{class:_(`${a(p)}__info`)},[u(a(ne),{src:(t=I.value)==null?void 0:t.banner,class:_(`${a(p)}__banner`)},null,8,["src","class"]),x("p",{class:_(`${a(p)}__brief`)},F(($=I.value)==null?void 0:$.brief),3),x("p",{class:_(`${a(p)}__usage`)},F((y=I.value)==null?void 0:y.usage),3)],2)]}),_:1},8,["visible","title"])]}),pluginStatus:s(({text:e})=>[e?(g(),v("span",fe,[u(a(w),{color:"green"},{default:s(()=>[ve]),_:1})])):(g(),v("span",Ce,[u(a(w),{color:"red"},{default:s(()=>[ye]),_:1})]))]),responseStatus:s(({text:e})=>[e?(g(),v("span",ke,[u(a(w),{color:"green"},{default:s(()=>[Se]),_:1})])):(g(),v("span",he,[u(a(w),{color:"red"},{default:s(()=>[be]),_:1})]))]),action:s(({record:e})=>[x("span",null,[u(a(S),{type:"primary",onClick:l=>B(e.gid,e.name,1),class:_(`${a(p)}__action`)},{default:s(()=>[De]),_:2},1032,["onClick","class"]),u(a(S),{danger:"",onClick:l=>B(e.gid,e.name,0),class:_(`${a(p)}__action`)},{default:s(()=>[Ie]),_:2},1032,["onClick","class"]),u(a(S),{type:"primary",onClick:l=>B(e.gid,e.name,2),class:_(`${a(p)}__action`)},{default:s(()=>[Fe]),_:2},1032,["onClick","class"]),u(a(S),{type:"primary",onClick:l=>N(e.gid,1),class:_(`${a(p)}__action`)},{default:s(()=>[xe]),_:2},1032,["onClick","class"]),u(a(S),{danger:"",onClick:l=>N(e.gid,0),class:_(`${a(p)}__action`)},{default:s(()=>[Le]),_:2},1032,["onClick","class"])])]),_:1},8,["dataSource"])]))}});export{Re as default};
