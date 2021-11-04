import './App.css';
import React, { Component } from 'react';
import 'antd/dist/antd.css';
import axios from 'axios';
import { message, Descriptions, Modal, Button, Table, Divider, Input, Form } from 'antd';
let qs = require('qs');

const api = "http://localhost:3000/";

const info = (msg) => {
  message.info(msg);
};

class List extends Component {
  constructor(props) {
    super(props);
    this.baseUrl = api;
    this.columns = [
      {
        title: 'platform',
        dataIndex: 'platform',
        key: 'platform',
      },
      {
        title: 'device_id_list',
        dataIndex: 'device_id_list',
        key: 'device_id_list',
      },
      {
        title: 'max_update_version_code',
        dataIndex: 'max_update_version_code',
        key: 'max_update_version_code',
      }, {
        title: 'min_update_version_code',
        dataIndex: 'min_update_version_code',
        key: 'min_update_version_code',
      }, {
        title: 'max_os_api',
        dataIndex: 'max_os_api',
        key: 'max_os_api',
      }, {
        title: 'min_os_api',
        dataIndex: 'min_os_api',
        key: 'min_os_api',
      }, {
        title: 'cpu_arch',
        dataIndex: 'cpu_arch',
        key: 'cpu_arch',
      }, {
        title: 'channel',
        dataIndex: 'channel',
        key: 'channel',
      }, {
        title: 'md5',
        dataIndex: 'md5',
        key: 'md5',
      }, {
        title: 'update_version_code',
        dataIndex: 'update_version_code',
        key: 'update_version_code',
      }, {
        title: 'is_available',
        dataIndex: 'is_available',
        key: 'is_available',
        render: text => <div>{String(text)}</div>,
      },
      {
        title: 'Action',
        key: 'action',
        render: (text, record) => (
          <span>
            <Button type="link" onClick={this.showMoreInfoModal.bind(this, record.title, record.update_tips, record.download_url)}>
              更多信息
            </Button>
            <Divider type="vertical" />
            <Button type="link" onClick={this.mydelete.bind(this, record.ID)}>
              删除
            </Button>
            <Divider type="vertical" />

            <Button type="link" onClick={this.mydisable.bind(this, record.ID)}>
              disable
            </Button>
            <Divider type="vertical" />

            <Button type="link" onClick={this.myenable.bind(this, record.ID)}>
              enable
            </Button>
          </span>
        ),
      },
    ];
    this.state = {
      list: [],
      visibleMoreInfo: false,
      ID: 0,
      title: "",
      update_tips: "",
      download_url: "",
      visible: false,
      confirmLoading: false,
    };
  }

  mydelete(id) {
    axios.delete(this.baseUrl + "rule/" + id)
      .then(res => {
        info(res.data.message);
        this.componentDidMount();
      })
      .catch(err => {
        console.log(err);
      });
  }
  mydisable(id) {
    axios.post(this.baseUrl + "disable-rule", qs.stringify({ ID: id }))
      .then(res => {
        info(res.data.message);
        this.componentDidMount();
      })
      .catch(err => {
        console.log(err);
      });
  }
  myenable(id) {
    axios.post(this.baseUrl + "enable-rule", qs.stringify({ ID: id }))
      .then(res => {
        info(res.data.message);
        this.componentDidMount();
      })
      .catch(err => {
        console.log(err);
      });
  }

  showMoreInfoModal = (title, update_tips, download_url) => {
    this.setState({
      visibleMoreInfo: true,
      title: title,
      update_tips: update_tips,
      download_url: download_url
    });
  };

  handleMoreInfoCancel = () => {
    this.setState({ visibleMoreInfo: false });
  };

  showModal = () => {
    this.setState({
      visible: true,
    });
  };

  handleCancel = () => {
    this.setState({
      visible: false,
    });
  };

  onFinish = (values) => {
    this.setState({
      confirmLoading: true,
    });
    axios.post(this.baseUrl + "rule", qs.stringify(values))
      .then(res => {
        info(res.data.message);
        this.componentDidMount();
        this.setState({
          visible: false,
          confirmLoading: false,
        });
      })
      .catch(err => {
        console.log(err);
      });
  };

  onFinishFailed = (errorInfo) => {
    info("Fail! 请填写完毕!");
  };

  render() {
    const { list, visibleMoreInfo, title, update_tips, download_url } = this.state;
    const { visible, confirmLoading, } = this.state;
    return (
      <div className="wrap">
        <div className="add">
          <Button type="primary" onClick={this.showModal} size="large">
            增加规则
          </Button>
          <Modal
            title="增加规则"
            visible={visible}
            confirmLoading={confirmLoading}
            onCancel={this.handleCancel}
            width="70%"
            footer={[
              <Button key="back" onClick={this.handleCancel}>
                退出
              </Button>
            ]}
          >
            <Form
              name="rule"
              labelCol={{
                span: 8,
              }}
              wrapperCol={{
                span: 12,
              }}
              onFinish={this.onFinish}
              onFinishFailed={this.onFinishFailed}
              autoComplete="off"
              className="formwarp"
            >
              <Form.Item
                label="platform"
                name="platform"
                rules={[
                  {
                    required: true,
                    message: 'platform!',
                  },
                ]}
              >
                <Input />
              </Form.Item>

              <Form.Item
                label="new_device_id_list"
                name="new_device_id_list"
                rules={[
                  {
                    required: true,
                    message: 'new_device_id_list!',
                  },
                ]}
              >
                <Input />
              </Form.Item>

              <Form.Item
                label="deleted_device_id_list"
                name="deleted_device_id_list"
                rules={[
                  {
                    required: true,
                    message: 'deleted_device_id_list!',
                  },
                ]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                label="max_update_version_code"
                name="max_update_version_code"
                rules={[
                  {
                    required: true,
                    message: 'max_update_version_code!',
                  },
                ]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                label="min_update_version_code"
                name="min_update_version_code"
                rules={[
                  {
                    required: true,
                    message: 'min_update_version_code!',
                  },
                ]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                label="max_os_api"
                name="max_os_api"
                rules={[
                  {
                    required: true,
                    message: 'max_os_api!',
                  },
                ]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                label="min_os_api"
                name="min_os_api"
                rules={[
                  {
                    required: true,
                    message: 'min_os_api!',
                  },
                ]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                label="cpu_arch"
                name="cpu_arch"
                rules={[
                  {
                    required: true,
                    message: 'cpu_arch!',
                  },
                ]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                label="channel"
                name="channele"
                rules={[
                  {
                    required: true,
                    message: 'channel!',
                  },
                ]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                label="download_url"
                name="download_url"
                rules={[
                  {
                    required: true,
                    message: 'download_url!',
                  },
                ]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                label="md5"
                name="md5rname"
                rules={[
                  {
                    required: true,
                    message: 'md5!',
                  },
                ]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                label="title"
                name="titleame"
                rules={[
                  {
                    required: true,
                    message: 'title!',
                  },
                ]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                label="update_tips"
                name="update_tips"
                rules={[
                  {
                    required: true,
                    message: 'update_tips!',
                  },
                ]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                label="update_version_code"
                name="update_version_code"
                rules={[
                  {
                    required: true,
                    message: 'update_version_code!',
                  },
                ]}
              >
                <Input />
              </Form.Item>
              <Form.Item
                label="is_available"
                name="is_available"
                rules={[
                  {
                    required: true,
                    message: 'is_available!',
                  },
                ]}
              >
                <Input />
              </Form.Item>

              <Form.Item
                wrapperCol={{
                  offset: 8,
                  span: 16,
                }}
              >
                <Button type="primary" htmlType="submit">
                  Submit
                </Button>
              </Form.Item>
            </Form>
          </Modal>
        </div>
        <div className="show">
          <Table rowKey="ID" columns={this.columns} dataSource={list} size="small" />
          <Modal
            visible={visibleMoreInfo}
            title="更多信息"
            onCancel={this.handleMoreInfoCancel}
            footer={[
              <Button key="back" onClick={this.handleMoreInfoCancel}>
                确认
              </Button>,

            ]}
          >
            <Descriptions title="" bordered>
              <Descriptions.Item label="title" span={3}>{title}</Descriptions.Item>
              <Descriptions.Item label="update_tips" span={3}>{update_tips}</Descriptions.Item>
              <Descriptions.Item label="download_url" span={3}>{download_url}</Descriptions.Item>
            </Descriptions>

          </Modal>
        </div>
      </div>
    );
  }
  // Ajax请求放在componentDidMount生命周期内
  componentDidMount() {
    // 使用axios完成ajax数据请求
    axios.get(this.baseUrl + "rule")
      .then(res => {
        const body = res.data.body;
        this.setState({
          list: body
        });
      })
      .catch(err => {
        console.log(err);
      });
  }
}

function App() {

  return (

    <div className="App">
      <List />
    </div>
  );

}

export default App;
