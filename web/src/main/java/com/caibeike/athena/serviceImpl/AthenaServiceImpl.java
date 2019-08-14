package com.caibeike.athena.serviceImpl;

import com.caibeike.athena.model.Data;
import com.caibeike.athena.model.Domain;
import com.caibeike.athena.model.Experiment;
import com.caibeike.athena.model.Layer;
import com.caibeike.athena.service.AthenaService;
import com.caibeike.athena.service.RpcService;
import org.springframework.stereotype.Service;

/**
 * Athena实验平台所有Rpc服务实现
 */
@Service(value = "athenaService")
public class AthenaServiceImpl implements AthenaService {

    @Override
    @RpcService("Athena.Manager.InsertDomain")
    public String InsertDomain(Domain domain) {
        return domain.toString();
    }

    @Override
    @RpcService("Athena.Manager.UpdateDomain")
    public String UpdateDomain(Domain domain) {
        return domain.toString();
    }

    @Override
    @RpcService("Athena.Manager.DeleteDomain")
    public String DeleteDomain(Long domainId) {
        return "{domainId: " + domainId + "}";
    }

    @Override
    @RpcService("Athena.Manager.InsertLayerInDomain")
    public String InsertLayerInDomain(Data data) {
        return data.toString();
    }

    @Override
    @RpcService("Athena.Manager.DeleteLayerInDomain")
    public String DeleteLayerInDomain(Data data) {
        return data.toString();
    }

    @Override
    @RpcService("Athena.Manager.UpdateLayer")
    public String UpdateLayer(Layer layer) {
        return layer.toString();
    }

    @Override
    @RpcService("Athena.Manager.DeleteLayer")
    public String DeleteLayer(Long layerId) {
        return "{layerId: " + layerId + "}";
    }

    @Override
    @RpcService("Athena.Manager.InsertExperimentInLayer")
    public String InsertExperimentInLayer(Data data) {
        return data.toString();
    }

    @Override
    @RpcService("Athena.Manager.UpdateExperiment")
    public String UpdateExperiment(Experiment experiment) {
        return experiment.toString();
    }

    @Override
    @RpcService("Athena.Manager.DeleteExperiment")
    public String DeleteExperiment(Long expId) {
        return "{exptId: " + expId + "}";
    }
}
