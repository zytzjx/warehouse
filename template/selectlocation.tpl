<form id="selectForm" class="layui-form" style="margin: 10px;">
    <div class="layui-form-item">
        <div class=" layui-input-inline">
            <label class="layui-form-label">Choose&nbsp;User:</label>           
        </div>
        <div class="layui-input-inline">
            <select id="options" name="options" lay-verify="required">
                {{range .}}              
                    <option value="option{{.ID}}">{{.Location}}</option>               
                {{end}}
            </select>
        </div>
    </div>
    <div class="layui-form-item" style="margin: 20px;">
        <button type="submit" class="layui-btn layui-btn-radius layui-btn-center" lay-submit lay-filter="submitForm">Submit</button>
    </div>
</form>