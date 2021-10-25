<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">

<xsl:template match="node">
  <div class="node">
    <p class="node-data">
      ID: <xsl:value-of select="@id" /><br />
      Описание: <xsl:value-of select="description" />
    </p>

    <div class="child-node">
      <xsl:apply-templates select="node" />
    </div>
  </div>
</xsl:template>

<xsl:template match="/">
  <html>
    <head>
      <link rel="stylesheet" type="text/css" href="styles.css" />
    </head>
    <body>
      <h2>My cum collection</h2>

      <xsl:apply-templates select="list/node" />
    </body>
  </html>
</xsl:template>

</xsl:stylesheet>
