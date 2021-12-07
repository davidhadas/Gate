package queuegate

import (
	"net/http"
	"sort"
	"strings"

	"go.uber.org/zap"
)

func getSortedKeys(m map[string][]string) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

func GateHandler(logger *zap.SugaredLogger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer h.ServeHTTP(w, r)

		logger.Infof("GateHandler Version %s", version)
		var keys []string

		for name, values := range r.Header {
			// Loop over all values for the name.
			for _, value := range values {
				logger.Infof("%s: %s", name, value)
			}
		}
		/*
			userAgentVals := r.Header.Get("User-Agent")

			contentEncoding := r.Header.Get("content-encoding")
			transferEncoding := r.Header.Get("transfer-encoding")
			keepAlive := r.Header.Get("keep-alive")
			connection := r.Header.Get("Connection")
			xForwardedFor := r.Header.Get("x-forwarded-for")
			cacheControl := r.Header.Get("cache-control")
			via := r.Header.Get("via")

			logger.Info("DH> userAgentVals ", userAgentVals)
			logger.Info("DH> contentEncoding ", contentEncoding)
			logger.Info("DH> transferEncoding ", transferEncoding)
			logger.Info("DH> keepAlive ", keepAlive)
			logger.Info("DH> connection ", connection)
			logger.Info("DH> xForwardedFor ", xForwardedFor)
			logger.Info("DH> cacheControl ", cacheControl)
			logger.Info("DH> via ", via)
		*/
		//var d = new Date();
		//h := make(map[string]string)
		var acceptHeaderVals, contentHeaderVals, userAgentVals, allHeaderKeys, otherHeaderVals, allHeaderVals, cookieVals strings.Builder

		//fingerprints := make([]string, 0, 12)
		//markers := make([]float32, 0, 12)
		//integers := make([]int, 0, 12)
		//roundedMarkers := make([]float32, 0, 12)
		//histograms := make([][]int, 0, 12)

		// Create a sorted slice of all header keys
		keys = getSortedKeys(r.Header)

		// Construct data about header keys and header vals
		for _, k := range keys {
			val := ""
			vals, ok := r.Header[k]
			if ok {
				val = strings.Join(vals, " ")
			}

			allHeaderKeys.WriteString(k)
			allHeaderVals.WriteString(val)
			switch {
			case strings.HasPrefix(k, "Accept"):
				acceptHeaderVals.WriteString(val)
			case strings.HasPrefix(k, "Content"):
				contentHeaderVals.WriteString(val)
			case k == "User-Agent":
				userAgentVals.WriteString(val)
			case strings.HasPrefix(k, "Cookie"):
				cookieVals.WriteString(val)
			default:
				otherHeaderVals.WriteString(val)
			}
		}
		logger.Infof("allHeaderKeys: %s", allHeaderKeys.String())
		logger.Infof("allHeaderVals: %s", allHeaderVals.String())
		logger.Infof("acceptHeaderVals: %s", acceptHeaderVals.String())
		logger.Infof("contentHeaderVals: %s", contentHeaderVals.String())
		logger.Infof("userAgentVals: %s", userAgentVals.String())
		logger.Infof("cookieVals: %s", cookieVals.String())
		logger.Infof("otherHeaderVals: %s", otherHeaderVals.String())

		// Create a sorted slice of all query leys
		query := r.URL.Query()
		keys = getSortedKeys(query)

		// Construct data about query keys and query vals
		for _, k := range keys {
			val := ""
			vals, ok := r.Header[k]
			if ok {
				val = strings.Join(vals, " ")
			}
			logger.Infof("query key and val is %s: %s", k, val)
		}

		// RequestURI is the unmodified request-target as sent by the client to a server.
		requestURI := r.RequestURI
		logger.Infof("requestURI: %s, Host: %s, Method: %s, Proto: %s, RemoteAddr: %s, ContentLength %d", requestURI, r.Host, r.Method, r.Proto, r.RemoteAddr, r.ContentLength)
		/*
			uriSplits := strings.Split(requestURI, "?")
			queryFragment
			queryKeys := ""
			queryContent := ""

			//elmsplit
			if (len(pathSplits)>1) {
				fragSplits
				qsplits :=  pathSplits[1].split("&")
				for (let elm of qsplits) {
					elmsplit = elm.split("=")
					queryKeys += decodeURIComponent(elmsplit[0])
					queryContent += decodeURIComponent(elmsplit[1])
				}
			}
			fingerprints.push(httpinfo["method"])
			fingerprints.push(httpinfo["scheme"])
			fingerprints.push(pathSplits[0])
			fingerprints.push(queryKeys)
			fingerprints.push(headers['transfer-encoding']||"")
			fingerprints.push(headers['content-encoding']||"")
			fingerprints.push(headers['keep-alive']||"")
			fingerprints.push(headers['connection']||"")
			fingerprints.push(headers['x-forwarded-for']||"")
			fingerprints.push(headers['cache-control']||"")
			fingerprints.push(headers['via']||"")
			fingerprints.push(acceptHeaderVals)
			fingerprints.push(contentHeaderVals)
			fingerprints.push(userAgentVals)
			fingerprints.push(allHeaderKeys)
			fingerprints.push(httpinfo["protocol"]||"")
			console.log(fingerprints)

			roundedMarkers.push(d.getDay()/6)
			roundedMarkers.push(d.getHours()/23)
			console.log(roundedMarkers)

			console.log(httpreq.body)
			console.log(otherHeaderVals)
			console.log(queryContent)


			integers.push(parseInt(httpreq.size)) // Content-Length  - size of body
			integers.push(otherHeaderVals.length)
			integers.push(queryContent.length)
			integers.push(cookieVals.length)
			integers.push(pathSplits[0].length)
			integers.push(allHeaderVals.length)
			console.log(markers, integers)



			histograms.push(hist(httpreq.body))
			histograms.push(hist(otherHeaderVals))
			histograms.push(hist(queryContent))
			histograms.push(hist(cookieVals))
			histograms.push(hist(allHeaderVals))
			console.log(histograms)

			fingerprint_path= pathSplits[0]


			var triggerInstance = headers["x-request-id"]||uuid.v4()


			for (var key in fingerprints) {
				fingerprints[key] = crypto.createHash('md5').update(fingerprints[key]).digest("hex").substr(0,16)
			}

			const dataout = JSON.stringify({
						gateId:   gate
					, serviceId: unit
					, triggerInstance: triggerInstance
					, data: {
							fingerprints: fingerprints
						, markers: markers
						, integers: integers
						, roundedMarkers: roundedMarkers
						, histograms: histograms
					}
				});

			console.log(unit, dataout);
			postRequest("Path: "+fingerprint_path, "/eval", dataout, callback)


		*/

	})
}

/*
  function hist(str) {
    str = str.toLowerCase();
    h = [0,0,0,0,0,0,0,0]
    for (i=0;i<str.length;i++) {
        c = str.charCodeAt(i)
        if ((c>=97 && c<=122) || (c>=48 && c<=57) || (c===32)) h[0]++
        else if (c>=127 || c<=31) h[1]++
        else if ((c===34) || (c===96) || (c===39)) h[2]++
        else if (c===59) h[3]++
        else h[7]++
    }
    h[4] = "count the number of  qoute start, qoute end, --, []"
    h[5] = (str.match(/0x/g) || []).length;
    h[6] = (str.match(/select|delete|drop|from|where/g) || []).length;
    console.log(h)
    return h
}


function Check(call, callback) {
*/
