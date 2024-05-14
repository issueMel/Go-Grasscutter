-- create protocol
kcp_protocol = Proto("KCP", "KCP Protocol")

-- fields for kcp  -- 28 bit
conv = ProtoField.uint64("kcp.conv", "conv")        -- 8
cmd = ProtoField.uint8("kcp.cmd", "cmd", base.DEC)  -- 1
frg = ProtoField.uint8("kcp.frg", "frg", base.DEC)  -- 1
wnd = ProtoField.uint16("kcp.wnd", "wnd", base.DEC) -- 2
ts = ProtoField.uint32("kcp.ts", "ts", base.DEC)    -- 4
sn = ProtoField.uint32("kcp.sn", "sn", base.DEC)    -- 4
una = ProtoField.uint32("kcp.una", "una", base.DEC) -- 4
len = ProtoField.uint32("kcp.len", "len", base.DEC) -- 4

kcp_protocol.fields = { conv, cmd, frg, wnd, ts, sn, una, len }

-- dissect each udp packet
function kcp_protocol.dissector(buffer, pinfo, tree)
    length = buffer:len()
    if length <= 20 then
        -- dissect packet in udp
        local cmd_id = buffer(0, 4)
        local b = buffer(4, 4)
        local c = buffer(8, 4)
        local d = buffer(12, 4)
        local e = buffer(16, 4)

        local first_cmd_name = get_packet_name(cmd_id:int())
        local info = string.format("[%s]", first_cmd_name)
        pinfo.cols.protocol = kcp_protocol.name
        udp_info = string.gsub(tostring(pinfo.cols.info), "Len", "Udplen", 1)
        pinfo.cols.info = string.gsub(udp_info, " U", info .. " U", 1)

        local tree_title = string.format(
                "KCP Protocol, CmdId = %d, B = %d, C = %d, D = %d, E = %d",
                cmd_id:int(),
                b:le_int(),
                c:le_int(),
                d:int(),
                e:int()
        )
        tree:add(kcp_protocol, buffer(), tree_title)
        return
    end

    local offset_s = 0
    local first_sn = buffer(offset_s + 16, 4):le_int()
    local first_len = buffer(offset_s + 24, 4):le_int()
    local first_cmd_name = get_cmd_name(buffer(offset_s + 8, 1):le_int())
    local info = string.format("[%s] Sn=%d Kcplen=%d", first_cmd_name, first_sn, first_len)

    pinfo.cols.protocol = kcp_protocol.name
    udp_info = string.gsub(tostring(pinfo.cols.info), "Len", "Udplen", 1)
    pinfo.cols.info = string.gsub(udp_info, " U", info .. " U", 1)

    -- dissect multi kcp packet in udp
    local offset = 0
    while offset < buffer:len() do
        local conv_buf = buffer(offset + 0, 8)
        local cmd_buf = buffer(offset + 8, 1)
        local wnd_buf = buffer(offset + 10, 2)
        local sn_buf = buffer(offset + 16, 4)
        local len_buf = buffer(offset + 24, 4)

        local cmd_name = get_cmd_name(cmd_buf:le_int())
        local data_len = len_buf:le_int()

        local tree_title = string.format(
                "KCP Protocol, %s, Sn: %d, Conv: %s, Wnd: %d, Len: %d",
                cmd_name,
                sn_buf:le_int(),
                conv_buf:le_int64(),
                wnd_buf:le_int(),
                data_len
        )
        local subtree = tree:add(kcp_protocol, buffer(), tree_title)
        subtree:add_le(conv, conv_buf)
        subtree:add_le(cmd, cmd_buf):append_text(" (" .. cmd_name .. ")")
        subtree:add_le(frg, buffer(offset + 9, 1))
        subtree:add_le(wnd, wnd_buf)
        subtree:add_le(ts, buffer(offset + 12, 4))
        subtree:add_le(sn, sn_buf)
        subtree:add_le(una, buffer(offset + 20, 4))
        subtree:add_le(len, len_buf)

        offset = offset + 28 + data_len
    end
end

function get_cmd_name(cmd_val)
    if cmd_val == 81 then
        return "PSH"
    elseif cmd_val == 82 then
        return "ACK"
    elseif cmd_val == 83 then
        return "ASK"
    elseif cmd_val == 84 then
        return "TELL"
    else
        return "UN_KNOW"
    end
end

function get_packet_name(val)
    if val == 255 then
        return "ASK"
    elseif val == 325 then
        return "TELL"
    elseif val == 404 then
        return "Disconnect"
    end
end

-- register kcp dissector to udp
local udp_port = DissectorTable.get("udp.port")
-- replace 8081 to the port for kcp
udp_port:add(22102, kcp_protocol)