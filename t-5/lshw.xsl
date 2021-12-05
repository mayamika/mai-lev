<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">

<xsl:template name="key-value">
  <xsl:param name="key" select="key" />
  <xsl:param name="value" select="value" />

  <xsl:choose>
    <xsl:when test="$value">
      <tr>
        <td class="key"><xsl:value-of select="$key" /></td>
        <td class="value"><xsl:value-of select="$value" /></td>
      </tr>
    </xsl:when>
  </xsl:choose>
</xsl:template>

<xsl:template match="node">
  <div class="node">
    <table class="node-data">
      <thead>
        <tr>
          <td class="key">ID:</td>
          <td class="value"><xsl:value-of select="@id" /></td>
        </tr>
      </thead>
      <tbody>
        <xsl:call-template name="key-value">
          <xsl:with-param name="key">
            Название:
          </xsl:with-param>
          <xsl:with-param name="value" select="product" />
        </xsl:call-template>

        <xsl:call-template name="key-value">
          <xsl:with-param name="key">
            Изготовитель:
          </xsl:with-param>
          <xsl:with-param name="value" select="vendor" />
        </xsl:call-template>

        <xsl:call-template name="key-value">
          <xsl:with-param name="key">
            Описание:
          </xsl:with-param>
          <xsl:with-param name="value" select="description" />
        </xsl:call-template>

      </tbody>
    </table>

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
      <h2 class="title">Список устройств</h2>

      <xsl:apply-templates select="node" />
    </body>
  </html>
</xsl:template>

</xsl:stylesheet>
