<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Device Info</title>
    <meta name="renderer" content="webkit">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <meta name="apple-mobile-web-app-status-bar-style" content="black"> 
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="format-detection" content="telephone=no">

    <script src="https://code.jquery.com/jquery-3.7.1.min.js" integrity="sha256-/JqT3SQfawRcv/BIHPThkBvs0OEvtFFmqPF/lYI/Cxo=" crossorigin="anonymous"></script>
    <link href="https://unpkg.com/layui@2.9.13/dist/css/layui.css" rel="stylesheet">
    <script src="https://unpkg.com/layui@2.9.13/dist/layui.js"></script> 

    <style>
      body{padding: 16px; padding-right: 32px;}
    </style>
  </head>
<body>
  <form class="layui-form layui-form-pane1" action="" lay-filter="first">
    <div class="layui-form-item">
      <label class="layui-form-label">Maker</label>
      <div class="layui-input-block">
        <select id="maker" name="maker" lay-filter="maker" lay-search="">
            {{range .Makers}}              
                <option value="{{.ID}}">{{.Maker}}</option>               
            {{end}}
        </select>
      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label">Model</label>
      <div class="layui-input-block">
        <input type="text" name="model" lay-verify="required|model" autocomplete="off" class="layui-input">
      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label">BarCode</label>
      <div class="layui-input-block">
        <input type="text" name="barcode" lay-verify="required|barcode" autocomplete="off" class="layui-input">
      </div>
    </div>
	<div class="layui-form-item">
      <label class="layui-form-label">Location</label>
      <div class="layui-input-block">
        <input type="text" name="location" lay-verify="required|location" autocomplete="off" class="layui-input">
      </div>
    </div>
	<div class="layui-form-item">
      <label class="layui-form-label">ESN</label>
      <div class="layui-input-block">
        <input type="text" name="esn" lay-verify="required|esn" autocomplete="off" class="layui-input">
      </div>
    </div>
	<div class="layui-form-item">
      <label class="layui-form-label">Borrower</label>
	  <div class="layui-input-block">
		  <select id="borrower"  name="borrower" lay-filter="borrower" lay-search="">
                {{range .Borrowers}}              
                    <option value="{{.ID}}">{{.Location}}</option>               
                {{end}}
		  </select>
	  </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label">FDModel</label>
      <div class="layui-input-block">
        <input type="text" name="fdmodel" autocomplete="off" class="layui-input">
      </div>
    </div>
    <div class="layui-form-item">
      <label class="layui-form-label">PhoneNumber</label>
      <div class="layui-input-block">
        <input type="text" name="phonenumber" autocomplete="off" class="layui-input">
      </div>
    </div>
	<div class="layui-form-item">
      <label class="layui-form-label">Carrier</label>
      <div class="layui-input-block">
        <input type="text" name="carrier" autocomplete="true" class="layui-input">
      </div>
    </div>
    
    <div class="layui-form-item layui-form-text">
      <label class="layui-form-label">Note</label>
      <div class="layui-input-block">
        <textarea placeholder="Please input Note" class="layui-textarea" name="note"></textarea>
      </div>
    </div>
    <!--div class="layui-form-item">
      <div class="layui-input-block">
        <button class="layui-btn" lay-submit lay-filter="update">Update</button>
        <button class="layui-btn" lay-submit  lay-filter="new">New</button>
		    <button class="layui-btn" >Cancel</button>
      </div>
    </div-->
  </form>

  <br><br><br>


  <script>
    function displayMessage (event){
      var param = event.data;

      let makerid=0;
      const selectMaker = document.getElementById("maker");
      for (var i = 0; i < selectMaker.options.length; i++) {
        if (selectMaker.options[i].label == param['maker']){
          makerid = parseInt(selectMaker.options[i].value);
          //$('#maker').val(selectMaker.options[i].value);         
          break;
        }
      }
      let borrowerid=0;
      const selectBorrower = document.getElementById("borrower");
      for (var i = 0; i < selectBorrower.options.length; i++) {
        if (selectBorrower.options[i].label == param['borrower']){
          borrowerid = parseInt(selectBorrower.options[i].value);                
          break;
        }
      }

			console.info('message1');
      layui.form.val('first', {
        'maker': makerid,
        'model': param['model']
        ,'barcode': param['barcode']
        ,'location': param['location']
        ,'borrower': borrowerid
        ,'esn': param['esn']
        ,'fdmodel': param['fdmodel']
        ,'phonenumber': param['phonenumber']
        ,'marketingname': param['marketingname']
        ,'carrier': param['carrier']
        ,'note': param['note']
        }); 
    }

    if (window.addEventListener) {
    // For standards-compliant web browsers
      console.info('message');
      window.addEventListener("message", displayMessage, false);
    }
    else {
      console.info('onmessage');
      window.attachEvent("onmessage", displayMessage);
    }


  layui.use(['form'], function(){
    var $ = layui.$;
    var form = layui.form;
    var layer = layui.layer;

    // 自定义验证规则
    form.verify({
      barcode: function(value){
        if(value && value.length < 5){ // 值若填写时才校验
          return 'Barcode length > 5';
        }
      },
      location: function(value) {
        if (value && value.length < 3) {
          return 'Location length > 3';
        }
      },
	  model: function(value) {
        if (value && value.length < 1) {
          return 'Model length > 1';
        }
      },
	  esn: function(value) {
        if (value && value.length < 1) {
          return 'ESN length > 1';
        }
      },
    });
    
    /*
    form.on('submit(top)', function(data){
      console.log(data);
      return false;
    });
    */
    
    //方法提交
  $('#testSubmit').on('click', function(){
    form.submit('top', function(data){
      layer.confirm('确定提交么？', function(index){
        layer.close(index);

        // 验证均通过后执行提交
        setTimeout(function(){
          alert(JSON.stringify(data.field));
        })
        
      });
    });
    return false;
  });
    
/*     
    //初始赋值
    var thisValue = form.val('first', {
      'maker': 1
      ,'model': 'xu@sentsin.com'
      ,'barcode': '2021-05-30'
      ,'location': '14H65'
      ,'borrower': 3
      ,'esn': 'ddddddddd'
      ,'fdmodel': 'fdsaf'
      ,'phonenumber': '+14082458880'
	  ,'marketingname':'mn_da'
	  ,'carrier':'Verizon'
      ,'note': 'phone test'
    });
*/  
    form.on('select(maker)', function(data){
      console.log('select.maker: ', this, data);
    });
	
    form.on('select(borrower)', function(data){
      console.log('select.borrower: ', this, data);
    });

    // 提交事件
    form.on('submit(update)', function(data){
	  //data.field['id']=123;
      //console.log(data)
      //alert(JSON.stringify(data.field));
	  //POST Data to server
      return false;
    });

    // 提交事件
    form.on('submit(new)', function(data){
      //console.log(data)
      alert(JSON.stringify(data.field));
	  //post data to server
      return false;
    });

  });

  </script>

</body>
</html>
