# fred_go_toolkit
FRED GoLang Toolkit for Federal Reserve Economic Data

[![Build Status](https://travis-ci.org/nswekosk/fred_go_toolkit.png?branch=master)](https://travis-ci.org/nswekosk/fred_go_toolkit)

[![GoDoc](https://godoc.org/github.com/nswekosk/fred_go_toolkit?status.svg)](https://godoc.org/github.com/nswekosk/fred_go_toolkit)

This is a GoLang toolkit for working with the Federal Reserve Economic Data. FRED contains frequently updated US macro and regional economic time series at annual, quarterly, monthly, weekly, and daily frequencies. FRED aggregates economic data from a variety of sources - most of which are US government agencies. The economic time series in FRED contains observation or measurement periods associated with data values. For instance, the US unemployment rate for the month of January, 1990 was 5.4 percent and for the month of January, 2000 was 4.0 percent.

## Installation

   Setup Go environment.
   Run 'go get github.com/nswekosk/fred_go_toolkit'

## Get a FRED API key

Sign up for a Fred API key: [http://api.stlouisfed.org/api_key.html](http://api.stlouisfed.org/api_key.html)

## Usage Example

fredConfig := FredConfig{
    APIKey: 'api-key',
    FileType: FileTypeJSON,     
    LogFile: "log.log",
}

client := CreateFredClient(fredConfig)  

    File types and log files are optional. You can use local constants 'FileTypeJSON' ('json')     
    or 'FileTypeXML' ('xml'). If no file type is specified, the default type is XML. 

params := make(map[string]interface{})

    These will be your input parameters for the API call. The value is of 
    type interface{} because it can take int's, string's, or booleans.    
    Each will be converted to string format for the url.                  

params["category_id"] = 125

    url: 'https://api.stlouisfed.org/fred/category?category_id=125&api_key=apiKey&file_type=json' 

fc, err := client.GetRecord(params)

    fc is of type FredClient. This contains all necessary attributes for all 
    responses for the available API calls.                                   


if err != nil {
  fmt.Println(err)
}

    NOTE: Logging is automatically generated once you run the client. The log  
    file will be created in the directory in which the client is instantiated. 
