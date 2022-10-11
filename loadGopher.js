var go2js = (() => {
    var c = Object.defineProperty;
    var f = r => c(r, "__esModule", {
        value: !0
    });
    var u = (r, o) => {
        f(r);
        for (var t in o) c(r, t, {
            get: o[t],
            enumerable: !0
        })
    };
    var y = {};
    u(y, {
        compile: () => g,
        format: () => a
    });
    
    var j = {},
        s = r => {
            if (r) {
                if (typeof r == "string") return r.endsWith("/") ? r.slice(0, -1) : r;
                throw new Error("invalid baseUrl")
            }
            let o = j?.url,
                t = "https://simonwaldherr.github.io/go2js/build/index.js", //document?.currentScript?.src
                e = "https://simonwaldherr.github.io/go2js/build/"; //location.href
            return (o && typeof o == "string" ? o : t && typeof t == "string" ? t : e && typeof e == "string" ? e : ".").split("/").slice(0, -1).join("/")
        };
    var g = async (r, o) => {
        let t = s(o);
        return await import("https://simonwaldherr.github.io/go2js/build/go2js-compile.js"), new Promise((e, n) => {
            globalThis.go2jsCompile(r, t, (i, m) => {
                i ? n(i) : e(m)
            })
        })
    };
    var a = async (r, o, t = !0) => (await import(s(o) + (t ? "/go2js-format.js" : "https://simonwaldherr.github.io/go2js/build/go2js-format-no-imports.js")), new Promise((i, m) => {
        let [p, l] = globalThis.go2jsFormat(r);
        l ? m(l) : i(p)
    }));
    return y;
})();
