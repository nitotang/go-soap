package soaphandler

var getTemplate = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" 
xmlns:blz="http://thomas-bayer.com/blz/">
   <soapenv:Header/>
   <soapenv:Body>
      <blz:getBank>
         <blz:blz>{{.Blz}}</blz:blz>
      </blz:getBank>
   </soapenv:Body>
</soapenv:Envelope> 
`
