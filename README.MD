# Invoice track

Parse XML files and obtain specific values from them.

## Description

Market Participants (MP) in the Wholesale Electrical Market in Mexico receive Statements of Account (ECD for its Spanish abbreviation) on a daily basis. These ECDs are issued by the Market Operator and serve as a guide to understand every single aspect of the MP’s operation.

## Instructions

The objective of this test is to write code that allows us to summarize the ECD contents for each MP every day. Specifically, we want to calculate the total balance for each invoice found in the ECD files.

The `<concepto>` nodes have a special child node `<monto_total>` that displays the monetary value represented by that node for the entire day. This is the value you need to extract from each _concepto_ node. For the sake of this test, you will only need to sum the `<monto_total>` inside the `<liquidacion num_liq="0">` invoices (`<factura>` nodes).

We have included a .NET project which can be used as a starter point. However, you can use another language if you wish.

In the attached code (_Program.cs_), we have provided some comments to guide you through the process:

1. **Change the location of the XML files** to a defined “xml repo” that should be located in the local Documents directory.
2. **Write the code to parse the XML files**. You can choose to either write local functions in Program.cs or develop the ECD_File class in Handler.cs to handle the processing.
3. **Define an appropriate structure** to return the necessary information. Keep in mind that the current code only iterates through a list of files and prints their names. We still need to extract data to aggregate the total for each invoice!
4. **Print the total of each invoice** for each provided ECD. You have flexibility in choosing the output format and style.
5. **Send your code and output file** to eduardo.hernandez@bravosenergia.com with the subject ‘ECD XML test - [your name initials]’. You can also share a git repo or a drive folder with the files. _Please also attach your professional resume in the email._

Finally, you may have noticed that a couple of comments start with ‘\*’. This indicates that it would be nice to make these parameters configurable from an external file (such as a config, JSON, or text file).

## Considerations

The current directory has the following items:

    TEC_TEST
    ├── ECD_Handler         -> Project ECD_Handler directory
    │   ├── bin                 -> Project binary files
    │   ├── dependencies        -> Directory for dependencies
    │   │   ├── Handler
    │   │   │   ├── Handler.cs          -> Handler dependency file (optional)
    │   ├── obj                 -> Project pre-compiled files
    │   ├── XML_files           -> Directory with the ECD files (XML files)
    │   ├── ECD_Handler.csproj  -> Project definition file
    │   ├── Program.cs          -> This is the main code file
    ├── README.md           -> This README file
    ├── tec_test.sln        -> Solution file

The ECD XML files have a specific structure that remains consistent. In summary, they follow this hierarchy:

    estadodecuenta
    ├── liquidaciones
    │   ├── liquidacion num_liq="0"
    │   │   ├── facturas
    │   │   │   ├── factura
    │   │   │   │   ├── conceptos
    │   │   │   │   │   ├── concepto
    │   │   │   │   │   │   ├── monto_total
    │   │   │   │   │   │   ├── ...
    │   ├── liquidacion num_liq="1"
    │   │   ...
