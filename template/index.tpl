<html lang="en">

<head>
    <link href="https://unpkg.com/layui@2.9.13/dist/css/layui.css" rel="stylesheet">
    <script src="https://unpkg.com/layui@2.9.13/dist/layui.js"></script>

    <link href="https://unpkg.com/tabulator-tables/dist/css/tabulator.min.css" rel="stylesheet">
    <script src="https://cdn.jsdelivr.net/npm/luxon/build/global/luxon.min.js"></script>
    <script type="text/javascript" src="https://unpkg.com/tabulator-tables/dist/js/tabulator.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/jsbarcode@3.11.0/dist/JsBarcode.all.min.js"></script>
    <!--xlsx js lib-->
    <script type="text/javascript" src="https://oss.sheetjs.com/sheetjs/xlsx.full.min.js"></script>
    <!--pdf js lib-->
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jspdf/2.4.0/jspdf.umd.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jspdf-autotable/3.5.20/jspdf.plugin.autotable.min.js"></script>
    <script src="https://code.jquery.com/jquery-3.7.1.min.js" integrity="sha256-/JqT3SQfawRcv/BIHPThkBvs0OEvtFFmqPF/lYI/Cxo=" crossorigin="anonymous"></script>
</head>

<body class="layui-padding-3">
    <h1>Barcode Generator</h1>
    <canvas id="barcode"></canvas>

    <script>
        // Your barcode data
        const barcodeData = "Designed by Jeffery";
        // Generate the barcode
        JsBarcode("#barcode", barcodeData, {
            format: "CODE128", // Barcode format
            lineColor: "#000", // Line color
            width: 2,          // Width of each barcode line
            height: 30,       // Height of the barcode
            displayValue: true // Display the value below the barcode
        });    
    </script>

    <div class="layui-btn-container">
        <button id="download-csv" type="button" class="layui-btn layui-btn-normal layui-btn-radius">Download
            CSV</button>
        <button id="download-json" type="button" class="layui-btn layui-btn-normal layui-btn-radius">Download
            JSON</button>
        <button id="download-xlsx" type="button" class="layui-btn layui-btn-normal layui-btn-radius">Download
            XLSX</button>
        <button id="download-pdf" type="button" class="layui-btn layui-btn-normal layui-btn-radius">Download
            PDF</button>
        <button id="download-html" type="button" class="layui-btn layui-btn-normal layui-btn-radius">Download
            HTML</button>
        <button id="print" type="button" class="layui-btn layui-btn-normal layui-btn-radius"><i
                class="layui-icon layui-icon-print"></i>print</button>
        <button id="logout" style="float:right" type="button" onclick="window.location.href = '/logout';"
                 class="layui-btn layui-btn-normal layui-btn-radius">Logout</button>
    </div>

    <div class="layui-form layui-form-item  layui-col-space32">
        <div class="layui-input-inline layui-col-md6">
            <select id="filter-field" lay-search="">
                <option></option>
                <option value="maker">Maker</option>
                <option value="model">Model</option>
                <option value="barcode">BarCode</option>
                <option value="location">Location</option>
                <option value="borrower">Borrower</option>
                <option value="esn">ESN</option>
                <option value="fdmodel">FD_Model</option>
                <option value="phonenumber">PhoneNumber</option>
                <option value="marketingname">MarketingName</option>
                <option value="carrier">Carrier</option>
                <option value="note">Note</option>
            </select>
        </div>
        <div class="layui-input-inline  layui-col-md6">
            <select id="filter-type" lay-search="cs">
                <option value="=">=</option>
                <option value="<">
                    < </option>
                <option value="<=">
                    <= </option>
                <option value=">">></option>
                <option value=">=">>=</option>
                <option value="!=">!=</option>
                <option value="like">like</option>
            </select>
        </div>
        <div class="layui-input-inline">
            <input class="layui-input" id="filter-value" type="text" placeholder="value to filter">
        </div>
        <div class="layui-input-inline">
            <button id="filter-clear" type="button" class="layui-btn layui-btn-normal layui-btn-radius">Clear
                Filter</button>
        </div>
        <div class="layui-input-inline">
            <label class="layui-form-label" style="padding-left: 400 px"><span id="search_count"></span> results in
                total <span id="total_count"></label>
        </div>
    </div>
    <div id='infocontent' style = "display : none">
           
    </div>

    <div id="devices-table"></div>
    <script type="text/javascript">

        function NewUpdateDialog(param){
            console.info(param)
            layui.use(['layer', 'form'], function () {
                var layer = layui.layer;
                var form = layui.form;
                console.info(param)
                
                layer.open({
                    type: 2,
                    title: 'Edit or Add',
                    content: '/editui',
                    area: ['500px', '750px'],
                    btn: ['Update', 'New', 'Cancel'],
                    yes: function(index, layero){
                       //Update
                       var iframeWindow = window[layero.find('iframe')[0]['name']];
                       var data = iframeWindow.layui.form.val('first',null);
                       data['id'] = param['id'];
                       data['maker'] = parseInt(data['maker']);
                       data['borrower'] = parseInt(data['borrower']);
                       jsonData = JSON.stringify(data);
                       console.info(jsonData);
                       jqxhr = $.post("/device", jsonData, (data, status) => {
                            console.log(data);

                        });
                        jqxhr.done(function(result) {
                            table.replaceData();
                        });
                        jqxhr.fail(function() {
                            alert("Return Device Request failed");
                        });
                    },
                    btn2: function(index, layero){
                        console.info("new")                       
                        //return false 开启该代码可禁止点击该按钮关闭
                        var iframeWindow = window[layero.find('iframe')[0]['name']];
                       var data = iframeWindow.layui.form.val('first',null);
                       //data['id'] = param['id'];
                       jsonData = JSON.stringify(data);
                       jqxhr = $.post("/device", jsonData, (data, status) => {
                            console.log(data);

                        });
                        jqxhr.done(function(result) {
                            table.replaceData();
                        });
                        jqxhr.fail(function() {
                            alert("Return Device Request failed");
                        });
                    },
                    btn3: function(index, layero){
                        console.info("close "+index.toString());
                        layer.confirm('Are you sure?', {icon: 3, title:'Infomation'}, function(indexa) { //只有当点击confirm框的确定时，该层才会关闭
                            console.info("dialog close "+indexa.toString());
                            layer.close(indexa)//self
                            layer.close(index) //parent
                        });
                        return false; 
                    },
                    success: function (layero, index) {   
                        var iframeWindow = window[layero.find('iframe')[0]['name']];
                        iframeWindow.postMessage(param, '*');
                        console.info('data sent');
                        /*
                        layero.find('iframe').on('load', function() {
                            // Get the iframe document after the layer is fully loaded
                            var elemMaker = iframeWindow.$('#maker'); // 获得 iframe 中某个输入框元素
                            console.info(elemMaker);
                            elemMaker.find('option').each(function() {
                                console.log('Value: ' + $(this).val() + ', Text: ' + $(this).text());
                                if($(this).text()==param['maker']){
                                    elemMaker.val($(this).val);
                                }
                            });  
                            iframeWindow.layui.form.render();
                        });       
                        iframeWindow.layui.form.val('first', {
                            'model': param['model']
                            ,'barcode': param['barcode']
                            ,'location': param['location']
                            ,'borrower': 3
                            ,'esn': param['esn']
                            ,'fdmodel': param['fdmodel']
                            ,'phonenumber': param['phonenumber']
                            ,'marketingname': param['marketingname']
                            ,'carrier': param['carrier']
                            ,'note': param['note']
                            });
                        */
                        form.render(); // Render the form
                    }
                });

                form.on('submit(submitForm)', function (data) {
                    var selectedOption = data.field.options;
                    layer.msg('You selected: ' + selectedOption);
                    console.info('You selected: ' + selectedOption + '   '+ data.field.value);
                    console.info(param);
                    layer.closeAll(); 
                    return false; 
                });
            });
        }

        function ShowLoctionDialog(param){
            console.info(param)
            layui.use(['layer', 'form'], function (param) {
                var layer = layui.layer;
                var form = layui.form;
                console.info(param)
                layer.open({
                    type: 1,
                    title: 'Select Location',
                    content:  `<form id="selectForm" class="layui-form" style="margin: 10px;">
                        <div class="layui-form-item">
                            <div class=" layui-input-inline">
                                <label class="layui-form-label">Choose&nbsp;User:</label>           
                            </div>
                            <div class="layui-input-inline">
                                <select id="options" lay-search="" lay-creatable="" name="options">
                                    {{range .}}              
                                        <option value="{{.ID}}">{{.Location}}</option>               
                                    {{end}}
                                </select>
                            </div>
                        </div>
                        <div class="layui-form-item" style="margin: 20px;">
                            <button type="submit" class="layui-btn layui-btn-radius layui-btn-center" lay-submit lay-filter="submitForm">Submit</button>
                        </div>
                    </form>`,
                    area: ['500px', '350px'],
                    success: function (layero, index) {
                        form.render(); // Render the form
                    }
                });

                form.on('submit(submitForm)', function (data) {
                    var selectedOption = data.field.options;
                    layer.msg('You selected: ' + selectedOption);
                    console.info('You selected: ' + selectedOption);
                    console.info(param);
                    const body = Array.from(table.getSelectedRows(), (x) => x.getData().id);
                        console.log(body);
                        let req = {
                            "borrower": selectedOption,
                            "ids": body
                        }
                        const jsonData = JSON.stringify(req);
                        console.log(jsonData);

                        jqxhr = $.post("/changeborrower", jsonData, (data, status) => {
                            console.log(data);

                        });
                        jqxhr.done(function(result) {
                            table.replaceData();
                        });
                        jqxhr.fail(function() {
                            alert("Return Device Request failed");
                        });

                    layer.closeAll(); 
                    return false; 
                });
            });
        }

        var rowMenu = [
            {
                label: "<i class='fas fa-user'></i>Change Location",
                menu:[
                    {
                        label: "<i class='fas fa-user'></i> Lend Device",
                        action: function (e, row) {
                            if (table.getSelectedRows().length ==0){
                                layer.alert('Please select device for returning.');
                                return
                            }
                            const body = Array.from(table.getSelectedRows(), (x) => x.getData().id);
                            console.log(body);
                            ShowLoctionDialog(body)
                        }
                    },
                    {
                        label: "<i class='fas fa-user'></i> Return Device",
                        action: function (e, row) {
                            if (table.getSelectedRows().length ==0){
                                layer.alert('Please select device for returning.');
                                return
                            }
                            const body = Array.from(table.getSelectedRows(), (x) => x.getData().id);
                            console.log(body);
                            let req = {
                                "user":"",
                                "ids": body
                            }
                            const jsonData = JSON.stringify(req);
                            console.log(jsonData);

                           jqxhr = $.post("/returnwarehouse", jsonData, (data, status) => {
                                    console.log(data);

                            });
                            jqxhr.done(function(result) {
                                 alert("POST success");
                            });
                            jqxhr.fail(function() {
                                alert("Return Device Request failed");
                            });
                        }
                    },
                    {
                        label: "<i class='fas fa-check-square'></i> Return to Customer",
                        action: function (e, row) {
                            layui.use('layer', function(){
                                var layer = layui.layer;
                                layer.open({
                                    title: 'User Info',
                                    content: `Mkaer: ${row.getData().maker} <br> Model: ${row.getData().model} <br>`
                                });
                            });
                        }
                    },
                    {
                        label: "<i class='fas fa-check-square'></i> Device Broken",
                        action: function (e, row) {
                            layui.use('layer', function(){
                                var layer = layui.layer;
                                layer.open({
                                    title: 'User Info',
                                    content: `Mkaer: ${row.getData().maker} <br> Model: ${row.getData().model} <br>`
                                });
                            });
                        }
                    },
                    {
                        label: "<i class='fas fa-check-square'></i> Device Missing",
                        action: function (e, row) {
                            layui.use('layer', function(){
                                var layer = layui.layer;
                                layer.open({
                                    title: 'User Info',
                                    content: `Mkaer: ${row.getData().maker} <br> Model: ${row.getData().model} <br>`
                                });
                            });
                        }
                    },
                ]
            },
            
            {
                label: "<i class='fas fa-check-square'></i> Device edit",
                action: function (e, row) {
                    var data = row.getData();
                    NewUpdateDialog(data)
                }
            },
        ]

        var table = new Tabulator("#devices-table", {
            maxHeight:"80%", 
            //height: 800, // set height of table to enable virtual DOM
            rowHeight: 35, //set rows to 40px height
            //data: tabledata, //load initial data into table
            ajaxURL: "/devices", //ajax URL
            printHeader: "<h1>Future Dial Devices<h1>",
            printFooter: "<h2>Future Dial Devices<h2>",

            layout: "fitColumns", //fit columns to width of table (optional)
            rowContextMenu: rowMenu,
            columns: [ //Define Table Columns
                { formatter: "rownum", hozAlign: "center", width: 50, headerSort: false },
                {
                    title: 'Select <br/> All <br/> <input type="checkbox" class="select-all-row" aria-label="select all rows" />',
                    field: 'IsSelected',
                    formatter: function (cell, formatterParams, onRendered) {
                        return `<input type="checkbox" class="select-row" aria-label="select this row" ${cell.getValue() ? 'checked' : ''} />`;
                    },
                    width: 50,
                    headerSort: false,
                    headerFilter: false,
                    cssClass: 'text-center',
                    frozen: true,
                    tooltips: false,
                    resizable: false,
                    cellClick: function (e, cell) {
                        const element = cell.getElement();
                        const checkbox = element.querySelector('.select-row');

                        if (cell.getData().IsSelected) {
                            cell.getRow().deselect();
                        } else {
                            cell.getRow().select();
                        }

                        document.querySelector('.select-all-row').checked = table.getSelectedRows().length === table.getDataCount();
                        checkbox.checked = !cell.getData().IsSelected;
                        cell.getData().IsSelected = !cell.getData().IsSelected;
                    },
                    headerClick: function (e, column) {
                        var allNotSelected = table.getSelectedRows().length !== table.getDataCount();
                        if (allNotSelected) {
                            table.selectRow();
                        } else {
                            table.deselectRow();
                        }

                        document.querySelectorAll('.select-all-row,.select-row').forEach(checkBox => checkBox.checked = allNotSelected);
                        table.getRows().forEach(row => row.update({ "IsSelected": allNotSelected }));
                    },
                },
                { title: "ID", field: "id", visible:false},
                { title: "Maker", field: "maker", sorter: "string", width: 100, headerFilter:"input" },
                { title: "Model", field: "model", sorter: "string", width: 100, headerFilter:"input" },
                { title: "BarCode", field: "barcode", sorter: "string", width: 80 },
                { title: "Location", field: "location", sorter: "string", width: 150 },
                { title: "Borrower", field: "borrower", sorter: "string", width: 150, headerFilter:"input"},
                { title: "ESN", field: "esn", sorter: "string", width: 150 },
                { title: "FD_Model", field: "fdmodel", sorter: "string", width: 150 },
                { title: "PhoneNumber", field: "phonenumber", sorter: "string", width: 150 },
                { title: "MarketingName", field: "marketingname", sorter: "string", width: 150 },
                { title: "Carrier", field: "carrier", sorter: "string", width: 100 },
                { title: "Note", field: "note", sorter: "string",  formatter:"textarea", editor:"textarea"},
            ],
        });

        //Define variables for input elements
        var fieldEl = document.getElementById("filter-field");
        var typeEl = document.getElementById("filter-type");
        var valueEl = document.getElementById("filter-value");

        //Trigger setFilter function with correct parameters
        function updateFilter() {
            var filterVal = fieldEl.options[fieldEl.selectedIndex].value;
            var typeVal = typeEl.options[typeEl.selectedIndex].value;

            var filter = filterVal == "function" ? customFilter : filterVal;

            if (filterVal == "function") {
                typeEl.disabled = true;
                valueEl.disabled = true;
            } else {
                typeEl.disabled = false;
                valueEl.disabled = false;
            }

            if (filterVal) {
                table.setFilter(filter, typeVal, valueEl.value);
            }
        }

        //Update filters on value change
        document.getElementById("filter-field").addEventListener("change", updateFilter);
        document.getElementById("filter-type").addEventListener("change", updateFilter);
        document.getElementById("filter-value").addEventListener("keyup", updateFilter);

        //Clear filters on "Clear Filters" button click
        document.getElementById("filter-clear").addEventListener("click", function () {
            fieldEl.value = "";
            typeEl.value = "=";
            valueEl.value = "";

            table.clearFilter();
        });
        //trigger an alert message when the row is clicked
        table.on("rowClick", function (e, row) {
            var rowElement = row.getElement();
            const checkbox = rowElement.querySelector('.select-row');
            checkbox.checked = !checkbox.checked;
            if (checkbox.checked) {
                row.select();
            } else {
                row.deselect();
            }

        });
        //table.on("rowAdded", function (row) {
            //row - row component
        //});
        table.on("cellEdited", function(cell){
            //cell - cell component
            console.info(cell.getValue());
            let req = {
                "note": cell.getValue(),
                "id": cell.getRow().getData()['id']
            }
            const jsonData = JSON.stringify(req);
            console.log(jsonData);

            jqxhr = $.post("/updatenote", jsonData, (data, status) => {
                console.log(data);

            });
            jqxhr.done(function(result) {
                table.replaceData();
            });
            jqxhr.fail(function() {
                alert("update notes failed");
            });            
        });
        table.on("cellEditCancelled", function(cell){
                //cell - cell component
            var odlvalue = cell.getOldValue();
            cell.setValue(odlvalue,true);
        });
        //table.on("rowDeleted", function (row) {
            //row - row component
        //});
        table.on("rowDblClick", function(e, row){
            //e - the click event object
            //row - row component
            NewUpdateDialog(row.getData());
        });
        //subscribe to event
        table.on("dataLoaded", function (data) {
            //data has been loaded
            var el = document.getElementById("total_count");
            el.innerHTML = data.length;
        });
        table.on("dataFiltered", function (filters, rows) {
            var el = document.getElementById("search_count");
            el.innerHTML = rows.length;
        });
        //trigger download of data.csv file
        document.getElementById("download-csv").addEventListener("click", function () {
            table.download("csv", "data.csv");
        });

        //trigger download of data.json file
        document.getElementById("download-json").addEventListener("click", function () {
            table.download("json", "data.json");
        });

        //trigger download of data.xlsx file
        document.getElementById("download-xlsx").addEventListener("click", function () {
            table.download("xlsx", "data.xlsx", { sheetName: "My Data" });
        });

        //trigger download of data.pdf file
        document.getElementById("download-pdf").addEventListener("click", function () {
            table.download("pdf", "data.pdf", {
                orientation: "portrait", //set page orientation to portrait
                title: "Example Report", //add title to report
            });
        });

        //trigger download of data.html file
        document.getElementById("download-html").addEventListener("click", function () {
            table.download("html", "data.html", { style: true });
        });
        document.getElementById("print").addEventListener("click", function () {
            table.printHeader = "<h1>THIS IS MY COOL TABLE</h1>"; // set header content on printed table
            table.printFooter = "<h3>THANKS FOR LOOKING AT MY TABLE</h3>"; // set footer content on printed table
            table.print(true, true);
        });

    </script>

</body>

</html>