var j=Object.defineProperty,L=Object.defineProperties;var T=Object.getOwnPropertyDescriptors;var $=Object.getOwnPropertySymbols;var V=Object.prototype.hasOwnProperty,O=Object.prototype.propertyIsEnumerable;var P=(e,t,r)=>t in e?j(e,t,{enumerable:!0,configurable:!0,writable:!0,value:r}):e[t]=r,k=(e,t)=>{for(var r in t||(t={}))V.call(t,r)&&P(e,r,t[r]);if($)for(var r of $(t))O.call(t,r)&&P(e,r,t[r]);return e},B=(e,t)=>L(e,T(t));var w=(e,t,r)=>new Promise((h,f)=>{var g=u=>{try{d(r.next(u))}catch(c){f(c)}},p=u=>{try{d(r.throw(u))}catch(c){f(c)}},d=u=>u.done?h(u.value):Promise.resolve(u.value).then(g,p);d((r=r.apply(e,t)).next())});import{A as x,cc as A,r as z,a as G,al as J,a0 as I,B as b,D as E,w as K,a5 as M,a1 as D,ac as Q,J as S,ad as q,K as F}from"./vendor.35b582d5.js";/* empty css               */import{_ as H,I as U,p as W,a as X,Q as Y,k as Z,b as ee,R as te,L as ne,s as ae,b8 as re,i as se}from"./index.044d747f.js";const oe=x({name:"LayoutBreadcrumb",components:{Icon:U,[A.name]:A},props:{theme:W.oneOf(["dark","light"])},setup(){const e=z([]),{currentRoute:t}=G(),{prefixCls:r}=X("layout-breadcrumb"),{getShowBreadCrumbIcon:h}=Y(),f=Z(),{t:g}=ee();J(()=>w(this,null,function*(){var C,y,R;if(t.value.name===te)return;const s=yield ne(),n=t.value.matched,a=n==null?void 0:n[n.length-1];let o=t.value.path;a&&((C=a==null?void 0:a.meta)==null?void 0:C.currentActiveMenu)&&(o=a.meta.currentActiveMenu);const i=ae(s,o),m=s.filter(N=>N.path===i[0]),l=p(m,i);if(!l||l.length===0)return;const _=d(l);((y=t.value.meta)==null?void 0:y.currentActiveMenu)&&_.push(B(k({},t.value),{name:((R=t.value.meta)==null?void 0:R.title)||t.value.name})),e.value=_}));function p(s,n){const a=[];return s.forEach(o=>{var i,m;n.includes(o.path)&&a.push(B(k({},o),{name:((i=o.meta)==null?void 0:i.title)||o.name})),((m=o.children)==null?void 0:m.length)&&a.push(...p(o.children,n))}),a}function d(s){return re(s,n=>{const{meta:a,name:o}=n;if(!a)return!!o;const{title:i,hideBreadcrumb:m}=a;return!(!i||m)}).filter(n=>{var a;return!((a=n.meta)==null?void 0:a.hideBreadcrumb)})}function u(s,n,a){a==null||a.preventDefault();const{children:o,redirect:i,meta:m}=s;if((o==null?void 0:o.length)&&!i){a==null||a.stopPropagation();return}if(!(m==null?void 0:m.carryParam))if(i&&se(i))f(i);else{let l="";n.length===1?l=n[0]:l=`${n.slice(1).pop()||""}`,l=/^\//.test(l)?l:`/${l}`,f(l)}}function c(s,n){return s.indexOf(n)!==s.length-1}function v(s){var n;return s.icon||((n=s.meta)==null?void 0:n.icon)}return{routes:e,t:g,prefixCls:r,getIcon:v,getShowBreadCrumbIcon:h,handleClick:u,hasRedirect:c}}}),ce={key:1};function ie(e,t,r,h,f,g){const p=I("Icon"),d=I("router-link"),u=I("a-breadcrumb");return b(),E("div",{class:F([e.prefixCls,`${e.prefixCls}--${e.theme}`])},[K(u,{routes:e.routes},{itemRender:M(({route:c,routes:v,paths:s})=>[e.getShowBreadCrumbIcon&&e.getIcon(c)?(b(),D(p,{key:0,icon:e.getIcon(c)},null,8,["icon"])):Q("",!0),e.hasRedirect(v,c)?(b(),D(d,{key:2,to:"",onClick:n=>e.handleClick(c,s,n)},{default:M(()=>[q(S(e.t(c.name||c.meta.title)),1)]),_:2},1032,["onClick"])):(b(),E("span",ce,S(e.t(c.name||c.meta.title)),1))]),_:1},8,["routes"])],2)}var fe=H(oe,[["render",ie]]);export{fe as default};
