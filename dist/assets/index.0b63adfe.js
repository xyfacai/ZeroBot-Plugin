import{A as v,r as a,j as x,u as e,B as y,D as w,w as H,a5 as S,H as j,K as u,X as d,aO as z}from"./vendor.35b582d5.js";import{u as C}from"./useWindowSizeFn.3778c3be.js";import{p as R,a as B,_ as b}from"./index.044d747f.js";import{a as F}from"./useContentViewHeight.263f1c49.js";const L=["src"],V=v({props:{frameSrc:R.string.def("")},setup(p){const i=a(!0),m=a(50),o=a(window.innerHeight),r=a(),{headerHeightRef:g}=F(),{prefixCls:c}=B("iframe-page");C(l,150,{immediate:!0});const f=x(()=>({height:`${e(o)}px`}));function l(){const n=e(r);if(!n)return;const t=g.value;m.value=t,o.value=window.innerHeight-t;const s=document.documentElement.clientHeight-t;n.style.height=`${s}px`}function _(){i.value=!1,l()}return(n,t)=>(y(),w("div",{class:u(e(c)),style:d(e(f))},[H(e(z),{spinning:i.value,size:"large",style:d(e(f))},{default:S(()=>[j("iframe",{src:p.frameSrc,class:u(`${e(c)}__main`),ref:(s,h)=>{h.frameRef=s,r.value=s},onLoad:_},null,42,L)]),_:1},8,["spinning","style"])],6))}});var E=b(V,[["__scopeId","data-v-179381bf"]]);export{E as default};
